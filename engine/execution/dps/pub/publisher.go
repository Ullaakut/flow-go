package pub

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/rs/zerolog"

	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/pub"
	_ "go.nanomsg.org/mangos/v3/transport/ws" // register ws transport (TODO: Import wss instead)

	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/engine/execution/state"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/pathfinder"
	"github.com/onflow/flow-go/ledger/complete"
)

type Publisher struct {
	log    zerolog.Logger
	unit   *engine.Unit
	stopCh chan struct{}

	// The compression and decompression components are necessary to
	// keep the incoming/outgoing packet size as small as possible.
	compressor   *zstd.Encoder
	decompressor *zstd.Decoder

	// The socket on which live updates are published.
	sock mangos.Socket

	// notifyCh is used to receive notifications from external components
	// when the state of the network has changed. Notifications can contain
	// new blocks or trie updates. In the future, they will also carry
	// events.
	notifyCh <-chan Notification
}

func New(log zerolog.Logger, host string, port uint16, notifyCh <-chan Notification) (*Publisher, error) {
	// Initialize compression/decompression components.
	compressor, err := zstd.NewWriter(nil)
	if err != nil {
		return nil, fmt.Errorf("could not initialize compressor: %w", err)
	}

	decompressor, err := zstd.NewReader(nil)
	if err != nil {
		return nil, fmt.Errorf("could not initialize decompressor: %w", err)
	}

	// Create and bind publish socket.
	sock, err := pub.NewSocket()
	if err != nil {
		return nil, fmt.Errorf("could not create pub socket: %w", err)
	}

	// TODO: Switch to WSS + IP Whitelist (https://github.com/optakt/flow-dps/issues/97)
	listenAddr := fmt.Sprint("ws://", host, ":", port)
	log.Info().Str("listen_address", listenAddr).Msg("starting pub server")
	err = sock.Listen(listenAddr)
	if err != nil {
		return nil, fmt.Errorf("could not listen on pub socket: %w", err)
	}

	p := Publisher{
		log: log.With().Str("component", "dps-publisher").Logger(),
		unit:   engine.NewUnit(),
		stopCh: make(chan struct{}),

		compressor:   compressor,
		decompressor: decompressor,

		sock: sock,

		notifyCh: notifyCh,
	}

	return &p, nil
}

func (p *Publisher) Ready() <-chan struct{} {
	p.unit.Launch(p.run)
	return p.unit.Ready()
}

func (p *Publisher) Done() <-chan struct{} {
	return p.unit.Done(p.stop)
}

func (p *Publisher) stop() {
	close(p.stopCh)
}

func (p *Publisher) run() {
	for {
		select {
		case <-p.stopCh:
			_ = p.sock.Close()
			return
		case notification := <-p.notifyCh:

			// Build a message to publish based on what kind of notification was received.
			// Messages are prefixed with their topic, which allows subscribers to only
			// receive the relevant information.
			var message []byte
			var err error
			switch notification.typ {
			case notificationTypeTrieUpdate:
				p.log.Debug().Msg("received trie update notification")
				message, err = p.trieNotifToMsg(notification)

			case notificationTypeBlock:
				p.log.Debug().Msg("received block update notification")
				message, err = p.blockNotifToMsg(notification)

			default:
				p.log.Error().Msg("unknown notification received")
				continue
			}
			if err != nil {
				p.log.Error().Err(err).Msg("unable to build message from notification")
				continue
			}

			// Publish the message on the socket.
			if err = p.sock.Send(message); err != nil {
				p.log.Error().Err(err).Msg("failed to publish update message")
			}

			p.log.Info().Msg("published state update")
		}
	}
}

// trieNotifToMsg encodes the contents of a Trie Update notification and publishes it on the socket.
func (p *Publisher) trieNotifToMsg(notif Notification) ([]byte, error) {
	ids, values := notif.view.RegisterUpdates()
	var update ledger.TrieUpdate
	for idx := range ids {
		key := state.RegisterIDToKey(ids[idx])
		path, err := pathfinder.KeyToPath(key, complete.DefaultPathFinderVersion)
		if err != nil {
			return nil, fmt.Errorf("trie update notification contains invalid key: %w", err)
		}

		update.Paths = append(update.Paths, path)
		update.Payloads = append(update.Payloads, ledger.NewPayload(key, values[idx]))
	}
	update.RootHash = notif.rootHash

	payload, err := cbor.Marshal(update)
	if err != nil {
		return nil, fmt.Errorf("could not marshal trie update notification: %w", err)
	}

	compressed := p.compressor.EncodeAll(payload, nil)

	// Add topic header.
	message := append(TopicTrie, compressed...)

	return message, nil
}

// blockNotifToMsg encodes the contents of a Block notification and publishes it on the socket.
func (p *Publisher) blockNotifToMsg(notif Notification) ([]byte, error) {
	payload, err := cbor.Marshal(notif.block)
	if err != nil {
		return nil, fmt.Errorf("could not marshal block notification: %w", err)
	}

	compressed := p.compressor.EncodeAll(payload, nil)

	// Add topic header.
	message := append(TopicBlock, compressed...)

	return message, nil
}
