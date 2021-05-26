package sync

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/zstd"
	pwal "github.com/m4ksio/wal/wal"
	"github.com/rs/zerolog"

	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/rep"
	_ "go.nanomsg.org/mangos/v3/transport/tcp" // register tcp transport (TODO: Import tlstcp instead)

	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/complete/wal"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/storage"
)

type Synchronizer struct {
	log    zerolog.Logger
	unit   *engine.Unit
	stopCh chan struct{}

	// The compression and decompression components are necessary to
	// keep the incoming/outgoing packet size as small as possible.
	compressor   *zstd.Encoder
	decompressor *zstd.Decoder

	// The socket on which live updates are published.
	sock mangos.Socket

	// Allows getting blocks at any height.
	blocks storage.Blocks

	// Allows replaying updates in order.
	trieDir string
}

func New(log zerolog.Logger, host string, port uint16, blocks storage.Blocks, trieDir string) (*Synchronizer, error) {
	// Initialize compression/decompression components.
	compressor, err := zstd.NewWriter(nil)
	if err != nil {
		return nil, fmt.Errorf("could not initialize compressor: %w", err)
	}

	decompressor, err := zstd.NewReader(nil)
	if err != nil {
		return nil, fmt.Errorf("could not initialize decompressor: %w", err)
	}

	// Create and bind synchronization socket.
	sock, err := rep.NewSocket()
	if err != nil {
		return nil, fmt.Errorf("could not create rep socket: %w", err)
	}

	// TODO: Switch to TLS on top of TCP + IP Whitelist (https://github.com/optakt/flow-dps/issues/97)
	listenAddr := fmt.Sprint("tcp://", host, ":", port)
	log.Info().Str("listen_address", listenAddr).Msg("starting sync server")
	err = sock.Listen(listenAddr)
	if err != nil {
		return nil, fmt.Errorf("could not listen on rep socket: %w", err)
	}

	s := Synchronizer{
		log:    log.With().Str("component", "dps-synchronizer").Logger(),
		unit:   engine.NewUnit(),
		stopCh: make(chan struct{}),

		compressor:   compressor,
		decompressor: decompressor,

		sock: sock,

		blocks: blocks,
		trieDir: trieDir,
	}

	return &s, nil
}

func (s *Synchronizer) Ready() <-chan struct{} {
	s.unit.Launch(s.run)
	return s.unit.Ready()
}

func (s *Synchronizer) Done() <-chan struct{} {
	return s.unit.Done(s.stop)
}

func (s *Synchronizer) stop() {
	close(s.stopCh)
}

func (s *Synchronizer) run() {
	// Since `s.sock.Recv()` is a blocking call, it needs to be closed from a separate goroutine
	// in case the Server's `Stop()` method is called.
	go func() {
		<-s.stopCh
		_ = s.sock.Close()
	}()

	for {
		// Only stop if the error is due to the socket being closed.
		err := s.handleRequest()
		if errors.Is(err, protocol.ErrClosed) {
			return
		}
		if err != nil {
			s.log.Error().Err(err).Msg("sync error")
		}
	}
}

// handleRequest handles one request, and can provide multiple responses if it
// needs to split its response into multiple chunks.
func (s *Synchronizer) handleRequest() error {
	ctx, err := s.sock.OpenContext()
	if err != nil {
		return err
	}
	defer ctx.Close()

	// Blocking call until a synchronization request is received.
	// Stop the Synchronizer if we get an error from having closed the socket.
	// Retry if any other error happened.
	msg, err := ctx.Recv()
	if err != nil {
		return err
	}

	// Switch on the topic of the synchronization request.
	switch {
	case bytes.HasPrefix(msg, TopicTrie):
		s.log.Debug().Msg("Received trie sync request")

		// Remove topic from message to get payload.
		payload := bytes.TrimPrefix(msg, TopicTrie)

		// Parse request and get response data.
		trieUpdates, err := s.trieRangeFromPayload(payload)
		if err != nil {
			s.sendSyncError(err)
			return err
		}

		// Send response data back over the socket.
		err = s.sendSyncResp(ctx, trieUpdates)
		if err != nil {
			s.sendSyncError(ErrInternal)
			return err
		}
		s.log.Debug().Int("elements", len(trieUpdates)).Msg("Trie sync response sent successfully")

	case bytes.HasPrefix(msg, TopicBlock):
		s.log.Debug().Msg("Received block sync request")

		// Remove topic from message to get payload.
		payload := bytes.TrimPrefix(msg, TopicBlock)

		// Parse request and get response data.
		blocks, err := s.blockRangeFromPayload(payload)
		if err != nil {
			s.sendSyncError(err)
			return err
		}

		// Send response data back over the socket.
		err = s.sendSyncResp(ctx, blocks)
		if err != nil {
			s.sendSyncError(ErrInternal)
			return err
		}
		s.log.Debug().Int("elements", len(blocks)).Msg("Block sync response sent successfully")

	default:
		s.log.Warn().Msg("unknown sync request received, ignoring")
		s.sendSyncError(ErrBadRequest)
	}

	return nil
}

// trieRangeFromPayload decodes the payload from a sync request and returns the data from
// the requested trie update range.
func (s *Synchronizer) trieRangeFromPayload(payload []byte) ([]ledger.TrieUpdate, error) {
	var req struct {
		From ledger.RootHash
		To   ledger.RootHash
	}

	request, err := s.decompressor.DecodeAll(payload, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to decode trie sync request payload: %s: %w", err, ErrBadRequest)
	}

	err = cbor.Unmarshal(request, &req)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal trie sync request payload: %s: %w", err, ErrBadRequest)
	}

	s.log.Debug().Hex("from", req.From[:]).Hex("to", req.To[:]).Msg("received valid trie sync request")

	segments, err := pwal.NewSegmentsReader(s.trieDir)
	if err != nil {
		return nil, fmt.Errorf("unable to open trie directory: %s: %w", err, ErrInternal)
	}

	ledgerWAL := pwal.NewReader(segments)

	var result []ledger.TrieUpdate
	var fromReached bool
	for {
		// This part reads the next entry from the WAL, makes sure we didn't
		// encounter an error when reading or decoding and ensures that it's a
		// trie update.
		next := ledgerWAL.Next()
		err := ledgerWAL.Err()
		if !next && err != nil {
			return nil, fmt.Errorf("could not read next record: %v: %w", err, ErrInternal)
		}
		if !next {
			s.log.Debug().Msg("reached end of ledgerWAL")
			break
		}

		record := ledgerWAL.Record()
		operation, _, update, err := wal.Decode(record)
		if err != nil {
			return nil, fmt.Errorf("could not decode record: %v: %w", err, ErrInternal)
		}
		if operation != wal.WALUpdate {
			continue
		}

		if update.RootHash == req.From {
			s.log.Debug().Msg("reached from")
			fromReached = true
		}

		// When we reach the final root hash we are looking for, we can stop directly
		// without including it in the response, since it has already been received
		// by the DPS through the publish socket.
		if update.RootHash == req.To {
			s.log.Debug().Msg("reached to")
			break
		}

		// Ignore update if it comes before the lower bound of requested the range.
		if !fromReached {
			s.log.Debug().Msg("skipping unrelated commit")
			continue
		}

		s.log.Debug().Msg("found update")
		result = append(result, *update)
	}

	// Limit to sending a chunk of 20 updates. The DPS keeps sending sync requests until it has all of the updates
	// it needs.
	if result != nil && len(result) > 20 {
		result = result[:20]
	}

	s.log.Info().
		Int("trie_updates", len(result)).
		Hex("from", req.From[:]).
		Hex("to", req.To[:]).
		Msg("Successfully retrieved trie updates")
	return result, nil
}

// trieRangeFromPayload decodes the payload from a sync request and returns the data from
// the requested height range.
func (s *Synchronizer) blockRangeFromPayload(payload []byte) ([]*flow.Block, error) {
	var req struct {
		From uint64
		To   uint64
	}

	request, err := s.decompressor.DecodeAll(payload, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to decode block sync request payload: %s: %w", err, ErrBadRequest)
	}

	err = cbor.Unmarshal(request, &req)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal block sync request payload: %s: %w", err, ErrBadRequest)
	}

	s.log.Debug().Uint64("from", req.From).Uint64("to", req.To).Msg("received valid block sync request")



	// Limit to sending a chunk of 20 blocks. The DPS keeps sending sync requests until it has all of the blocks
	// it needs.
	to := req.To
	if req.To - req.From > 20 {
		to = req.From + 20
	}

	var blocks []*flow.Block
	for height := req.From; height <= to; height++ {
		block, err := s.blocks.ByHeight(height)
		if err != nil {
			return nil, fmt.Errorf("unable to find block at height %d: %s: %w", height, err, ErrInternal)
		}

		blocks = append(blocks, block)
	}

	s.log.Info().
		Int("blocks", len(blocks)).
		Uint64("from", req.From).
		Uint64("to", to).
		Msg("Successfully retrieved blocks")

	return blocks, nil
}

// sendSyncResp marshals the given response and sends it over the sync socket.
func (s *Synchronizer) sendSyncResp(ctx mangos.Context, resp interface{}) error {
	payload, err := cbor.Marshal(resp)
	if err != nil {
		return fmt.Errorf("unable to marshal sync response")
	}

	message := s.compressor.EncodeAll(payload, nil)

	return ctx.Send(message)
}

//// sendBlockSyncResp marshals the given blocks and sends them over the sync socket by chunks of 20.
//func (s *Synchronizer) sendBlockSyncResp(ctx mangos.Context, blocks []*flow.Block) error {
//	chunk := make([]*flow.Block, 0, 20)
//	for idx, block := range blocks {
//		chunk = append(chunk, block)
//
//		// Every 20 elements, send the chunk and reset it to be empty.
//		if idx % 20 == 0 {
//			s.log.Info().Msg("BEFORE SEND") // FIXME: Remove
//			err := s.sendSyncResp(ctx, chunk)
//			if err != nil {
//				return err
//			}
//			s.log.Info().Msg("AFTER SEND") // FIXME: Remove
//
//			chunk = make([]*flow.Block, 0, 20)
//		}
//	}
//
//	// If there are no left over blocks after chunk creation, stop here.
//	if len(chunk) == 0 {
//		return nil
//	}
//
//	// Otherwise, send the remaining blocks.
//	return s.sendSyncResp(ctx, chunk)
//}

//// sendTrieSyncResp marshals the given trie updates and sends them over the sync socket by chunks of 20.
//func (s *Synchronizer) sendTrieSyncResp(ctx mangos.Context, updates []ledger.TrieUpdate) error {
//	chunk := make([]ledger.TrieUpdate, 0, 20)
//	for idx, update := range updates {
//		chunk = append(chunk, update)
//
//		// Every 20 elements, send the chunk and reset it to be empty.
//		if idx % 20 == 0 {
//
//			err := s.sendSyncResp(ctx, chunk)
//			if err != nil {
//				return err
//			}
//
//			chunk = make([]ledger.TrieUpdate, 0, 20)
//		}
//	}
//
//	// If there are no left over blocks after chunk creation, stop here.
//	if len(chunk) == 0 {
//		return nil
//	}
//
//	// Otherwise, send the remaining blocks.
//	return s.sendSyncResp(ctx, chunk)
//}

// sendSyncError marshals the given error and sends it over the sync socket.
// It cannot fail.
func (s *Synchronizer) sendSyncError(err error) {
	payload, err := cbor.Marshal(err)
	if err != nil {
		s.log.Error().Err(err).Msg("unable to marshal error to send to sync client")
		return
	}

	message := s.compressor.EncodeAll(payload, nil)

	_ = s.sock.Send(message)
}
