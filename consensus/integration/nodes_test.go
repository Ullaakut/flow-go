package integration_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/dapperlabs/flow-go/cmd/bootstrap/run"
	"github.com/dapperlabs/flow-go/consensus"
	"github.com/dapperlabs/flow-go/consensus/hotstuff"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/committee"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/committee/leader"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/helper"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/notifications"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/notifications/pubsub"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/persister"
	synceng "github.com/dapperlabs/flow-go/engine/common/synchronization"
	"github.com/dapperlabs/flow-go/engine/consensus/compliance"
	"github.com/dapperlabs/flow-go/model/bootstrap"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/module/buffer"
	builder "github.com/dapperlabs/flow-go/module/builder/consensus"
	finalizer "github.com/dapperlabs/flow-go/module/finalizer/consensus"
	"github.com/dapperlabs/flow-go/module/local"
	"github.com/dapperlabs/flow-go/module/mempool/stdmap"
	"github.com/dapperlabs/flow-go/module/metrics"
	"github.com/dapperlabs/flow-go/module/signature"
	synccore "github.com/dapperlabs/flow-go/module/synchronization"
	"github.com/dapperlabs/flow-go/module/trace"
	networkmock "github.com/dapperlabs/flow-go/network/mock"
	protocol "github.com/dapperlabs/flow-go/state/protocol/badger"
	storage "github.com/dapperlabs/flow-go/storage/badger"
	storagemock "github.com/dapperlabs/flow-go/storage/mock"
	"github.com/dapperlabs/flow-go/utils/unittest"
)

const hotstuffTimeout = 2 * time.Second

type Node struct {
	db            *badger.DB
	dbDir         string
	index         int
	log           zerolog.Logger
	id            *flow.Identity
	compliance    *compliance.Engine
	sync          *synceng.Engine
	hot           *hotstuff.EventLoop
	state         *protocol.State
	headers       *storage.Headers
	net           *Network
	counter       *CounterConsumer
	blockproposal int
	blockvote     int
	syncreq       int
	syncresp      int
	rangereq      int
	batchreq      int
	batchresp     int
}

func createNodes(t *testing.T, n int, onStop func([]*Node) bool) ([]*Node, *Stopper, *Hub) {

	// create n consensus node participants
	consensus := unittest.IdentityListFixture(n, unittest.WithRole(flow.RoleConsensus))

	// create non-consensus nodes
	collection := unittest.IdentityFixture(unittest.WithRole(flow.RoleCollection))
	verification := unittest.IdentityFixture(unittest.WithRole(flow.RoleVerification))
	execution := unittest.IdentityFixture(unittest.WithRole(flow.RoleExecution))

	// append additional nodes to consensus
	participants := append(consensus, collection, verification, execution)

	chainID := flow.Testnet
	parentID := flow.ZeroID
	height := uint64(0)
	timestamp := time.Now().UTC()
	// add all identities to rootBlock block and
	// create and bootstrap consensus node with the rootBlock
	rootBlock := run.GenerateRootBlock(chainID, parentID, height, timestamp, participants)

	// make root QC
	sig1 := make([]byte, 32)
	rand.Read(sig1[:])
	sig2 := make([]byte, 32)
	rand.Read(sig2[:])
	c := &signature.Combiner{}
	combined, err := c.Join(sig1, sig2)
	require.NoError(t, err)

	// all participants will sign the rootBlock block
	signerIDs := make([]flow.Identifier, 0)
	// only consensus participants can sign root block
	for _, participant := range consensus {
		signerIDs = append(signerIDs, participant.ID())
	}

	rootQC := &model.QuorumCertificate{
		View:      rootBlock.Header.View,
		BlockID:   rootBlock.ID(),
		SignerIDs: signerIDs,
		SigData:   combined,
	}

	hub := NewHub()
	stopper := NewStopper(onStop)
	nodes := make([]*Node, 0, len(consensus))
	for i, identity := range consensus {
		node := createNode(t, i, identity, consensus, rootBlock, rootQC, hub, stopper)
		nodes = append(nodes, node)
	}

	return nodes, stopper, hub
}

func createNode(t *testing.T, index int, identity *flow.Identity, participants flow.IdentityList, rootBlock *flow.Block, rootQC *model.QuorumCertificate, hub *Hub, stopper *Stopper) *Node {
	db, dbDir := unittest.TempBadgerDB(t)

	metrics := metrics.NewNoopCollector()
	tracer := trace.NewNoopTracer()

	headersDB := storage.NewHeaders(metrics, db)
	identitiesDB := storage.NewIdentities(metrics, db)
	guaranteesDB := storage.NewGuarantees(metrics, db)
	sealsDB := storage.NewSeals(metrics, db)
	indexDB := storage.NewIndex(metrics, db)
	payloadsDB := storage.NewPayloads(db, indexDB, identitiesDB, guaranteesDB, sealsDB)
	blocksDB := storage.NewBlocks(db, headersDB, payloadsDB)

	state, err := protocol.NewState(metrics, db, headersDB, identitiesDB, sealsDB, indexDB, payloadsDB, blocksDB)
	require.NoError(t, err)

	result := bootstrap.Result(rootBlock, unittest.GenesisStateCommitment)
	seal := bootstrap.Seal(result)
	err = state.Mutate().Bootstrap(rootBlock, result, seal)
	require.NoError(t, err)

	localID := identity.ID()

	node := &Node{
		db:    db,
		dbDir: dbDir,
		index: index,
		id:    identity,
	}

	// log with node index an ID
	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	log := zerolog.New(os.Stderr).Level(zerolog.WarnLevel).With().Timestamp().Int("index", index).Hex("node_id", localID[:]).Logger()

	stopConsumer := stopper.AddNode(node)

	counterConsumer := &CounterConsumer{
		log:      log,
		interval: time.Second,
		next:     time.Now().Add(time.Second),
		finalized: func(total uint) {
			stopper.onFinalizedTotal(node.id.ID(), total)
		},
	}

	// log with node index
	notifier := notifications.NewLogConsumer(log)
	dis := pubsub.NewDistributor()
	dis.AddConsumer(stopConsumer)
	dis.AddConsumer(counterConsumer)
	dis.AddConsumer(notifier)

	cleaner := &storagemock.Cleaner{}
	cleaner.On("RunGC")

	// make local
	priv := helper.MakeBLSKey(t)
	local, err := local.New(identity, priv)
	require.NoError(t, err)

	// add a network for this node to the hub
	net := hub.AddNetwork(localID, node)

	guaranteeLimit, sealLimit := uint(1000), uint(1000)
	guarantees, err := stdmap.NewGuarantees(guaranteeLimit)
	require.NoError(t, err)
	seals, err := stdmap.NewSeals(sealLimit)
	require.NoError(t, err)

	// initialize the block builder
	build := builder.NewBuilder(metrics, db, headersDB, sealsDB, indexDB, blocksDB, guarantees, seals)

	signer := &Signer{identity.ID()}

	// initialize the pending blocks cache
	cache := buffer.NewPendingBlocks()

	rootHeader := rootBlock.Header

	// initialize and pre-generate leader selections from the seed
	selection, err := leader.NewSelectionForConsensus(10000, rootHeader, rootQC, state)
	require.NoError(t, err)

	// selector := filter.HasRole(flow.RoleConsensus)
	com, err := committee.NewMainConsensusCommitteeState(state, localID, selection)
	require.NoError(t, err)

	// initialize the block finalizer
	final := finalizer.NewFinalizer(db, headersDB, state)

	// initialize the persister
	persist := persister.New(db)

	prov := &networkmock.Engine{}
	prov.On("SubmitLocal", mock.Anything).Return(nil)

	syncCore, err := synccore.New(log, synccore.DefaultConfig())
	require.NoError(t, err)

	// initialize the compliance engine
	comp, err := compliance.New(log, metrics, tracer, metrics, metrics, net, local, cleaner, headersDB, payloadsDB, state, prov, cache, syncCore)
	require.NoError(t, err)

	// initialize the synchronization engine
	sync, err := synceng.New(log, metrics, net, local, state, blocksDB, comp, syncCore)
	require.NoError(t, err)

	pending := []*flow.Header{}
	// initialize the block finalizer
	hot, err := consensus.NewParticipant(log, tracer, dis, metrics, headersDB,
		com, build, final, persist, signer, comp, rootHeader,
		rootQC, rootHeader, pending, consensus.WithInitialTimeout(hotstuffTimeout))

	require.NoError(t, err)

	comp = comp.WithConsensus(hot)

	node.compliance = comp
	node.sync = sync
	node.state = state
	node.hot = hot
	node.headers = headersDB
	node.net = net
	node.log = log
	node.counter = counterConsumer

	return node
}

func cleanupNodes(nodes []*Node) {
	for _, n := range nodes {
		n.db.Close()
		os.RemoveAll(n.dbDir)
	}
}
