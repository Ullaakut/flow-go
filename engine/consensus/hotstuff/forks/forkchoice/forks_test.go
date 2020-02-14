package forkchoice

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/dapperlabs/flow-go/crypto"
	"github.com/dapperlabs/flow-go/engine/consensus/hotstuff"
	"github.com/dapperlabs/flow-go/engine/consensus/hotstuff/forks"
	"github.com/dapperlabs/flow-go/engine/consensus/hotstuff/forks/finalizer"
	mockdist "github.com/dapperlabs/flow-go/engine/consensus/hotstuff/notifications/mock"
	"github.com/dapperlabs/flow-go/engine/consensus/hotstuff/types"
	"github.com/dapperlabs/flow-go/model/flow"
	mockfinalizer "github.com/dapperlabs/flow-go/module/mock"
	"github.com/stretchr/testify/assert"
)

// TestForks_ImplementsInterface tests that forks.Forks implements hotstuff.Forks
// (compile-time test)
func TestForks_ImplementsInterface(t *testing.T) {
	var _ hotstuff.Forks = &forks.Forks{}
}

// TestForks_Initialization tests that Forks correctly reports trusted Root
func TestForks_Initialization(t *testing.T) {
	forks, _, _, root := initForks(t, 1)

	assert.Equal(t, forks.FinalizedView(), uint64(1))
	assert.Equal(t, forks.FinalizedBlock(), root.Block)

	assert.Equal(t, forks.GetBlocksForView(0), []*types.Block{})
	assert.Equal(t, forks.GetBlocksForView(1), []*types.Block{root.Block})
	assert.Equal(t, forks.GetBlocksForView(2), []*types.Block{})

	b, found := forks.GetBlock(root.Block.BlockID)
	assert.True(t, found, "Missing trusted Root ")
	assert.Equal(t, root.Block, b)
}

// TestForks_AddBlock verifies that Block can be added
func TestForks_AddBlock(t *testing.T) {
	forks, _, notifier, root := initForks(t, 1)

	block02 := makeBlock(2, root.QC, flow.ZeroID)
	notifier.On("OnBlockIncorporated", block02).Return().Once()
	err := forks.AddBlock(block02)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	notifier.AssertExpectations(t)

	assert.Equal(t, forks.GetBlocksForView(2), []*types.Block{block02})
	b, found := forks.GetBlock(block02.BlockID)
	assert.True(t, found)
	assert.Equal(t, block02, b)
}

// TestForks_3ChainFinalization tests happy-path direct 3-chain finalization
func TestForks_3ChainFinalization(t *testing.T) {
	forks, finCallback, notifier, root := initForks(t, 1) // includes genesis block (v1)

	block2 := makeBlock(2, root.QC, flow.ZeroID)
	notifier.On("OnBlockIncorporated", block2).Return().Once()
	addBlock2Forks(t, block2, forks)
	notifier.AssertExpectations(t)

	block3 := makeBlock(3, qc(block2.View, block2.BlockID), flow.ZeroID)
	notifier.On("OnBlockIncorporated", block3).Return().Once()
	notifier.On("OnQcIncorporated", block3.QC).Return().Once()
	addBlock2Forks(t, block3, forks)
	notifier.AssertExpectations(t)

	// creates direct 3-chain on genesis block (1), which is already finalized
	block4 := makeBlock(4, qc(block3.View, block3.BlockID), flow.ZeroID)
	notifier.On("OnBlockIncorporated", block4).Return().Once()
	notifier.On("OnQcIncorporated", block4.QC).Return().Once()
	addBlock2Forks(t, block4, forks)
	notifier.AssertExpectations(t)

	// creates direct 3-chain on block (2) => finalize (2)
	block5 := makeBlock(5, qc(block4.View, block4.BlockID), flow.ZeroID)
	notifier.On("OnBlockIncorporated", block5).Return().Once()
	notifier.On("OnQcIncorporated", block5.QC).Return().Once()
	notifier.On("OnFinalizedBlock", block2).Return().Once()
	finCallback.On("MakeFinal", block2.BlockID).Return(nil).Once()
	addBlock2Forks(t, block5, forks)
	notifier.AssertExpectations(t)
	finCallback.AssertExpectations(t)

	// creates direct 3-chain on block (3) => finalize (3)
	block6 := makeBlock(6, qc(block5.View, block5.BlockID), flow.ZeroID)
	notifier.On("OnBlockIncorporated", block6).Return().Once()
	notifier.On("OnQcIncorporated", block6.QC).Return().Once()
	notifier.On("OnFinalizedBlock", block3).Return().Once()
	finCallback.On("MakeFinal", block3.BlockID).Return(nil).Once()
	addBlock2Forks(t, block6, forks)
	notifier.AssertExpectations(t)
	finCallback.AssertExpectations(t)
}

func addBlock2Forks(t *testing.T, block *types.Block, forks hotstuff.Forks) {
	err := forks.AddBlock(block)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	verifyStored(t, block, forks)
}

// verifyStored verifies that block is stored in forks
func verifyStored(t *testing.T, block *types.Block, forks hotstuff.Forks) {
	b, found := forks.GetBlock(block.BlockID)
	assert.True(t, found)
	assert.Equal(t, block, b)

	found = false
	siblings := forks.GetBlocksForView(block.View)
	assert.True(t, len(siblings) > 0)
	for _, b := range forks.GetBlocksForView(block.View) {
		if b != block {
			continue
		}
		if found { // we already found block in slice, i.e. this is a duplicate
			assert.Fail(t, fmt.Sprintf("Duplicate block: %v", block.BlockID))
		}
		found = true
	}
	assert.True(t, found, fmt.Sprintf("Did not find block: %v", block.BlockID))
}

func initForks(t *testing.T, view uint64) (*forks.Forks, *mockfinalizer.Finalizer, *mockdist.Consumer, *forks.BlockQC) {
	notifier := &mockdist.Consumer{}
	finalizationCallback := &mockfinalizer.Finalizer{}

	// construct Finalizer
	root := makeRootBlock(t, view)
	notifier.On("OnBlockIncorporated", root.Block).Return().Once()
	fnlzr, err := finalizer.New(root, finalizationCallback, notifier)

	// construct ForkChoice
	notifier.On("OnQcIncorporated", root.QC).Return().Once()
	fc, err := NewNewestForkChoice(fnlzr, notifier)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	notifier.AssertExpectations(t)
	return forks.New(fnlzr, fc), finalizationCallback, notifier, root
}

func makeRootBlock(t *testing.T, view uint64) *forks.BlockQC {
	// construct Finalizer with Genesis Block
	genesisBlock := makeBlock(view, nil, flow.ZeroID)
	genesisQC := qc(view, genesisBlock.BlockID)
	root := forks.BlockQC{Block: genesisBlock, QC: genesisQC}
	return &root
}

func qc(view uint64, id flow.Identifier) *types.QuorumCertificate {
	return &types.QuorumCertificate{View: view, BlockID: id}
}

func makeChildBlock(blockView uint64, payloadHash flow.Identifier, parent *types.Block) *types.Block {
	qcForParent := qc(parent.View, parent.BlockID)
	id := computeID(blockView, qcForParent, payloadHash)
	return &types.Block{
		BlockID:     id,
		View:        blockView,
		QC:          qcForParent,
		PayloadHash: payloadHash,
		Height:      0,
		ChainID:     "",
	}
}

func makeBlock(blockView uint64, blockQC *types.QuorumCertificate, payloadHash flow.Identifier) *types.Block {
	if blockQC == nil {
		blockQC = qc(0, flow.Identifier{})
	}
	id := computeID(blockView, blockQC, payloadHash)
	return &types.Block{
		BlockID:     id,
		View:        blockView,
		QC:          blockQC,
		PayloadHash: payloadHash,
		Height:      0,
		ChainID:     "",
	}
}

func string2Identifyer(s string) flow.Identifier {
	var identifier flow.Identifier
	copy(identifier[:], []byte(s))
	return identifier
}

// computeID is an INCOMPLETE STUB needed so we can test Forks.
// When computing the Block's ID, this implementation only considers
// the fields used by Forks.
// TODO need full implementation
func computeID(view uint64, qc *types.QuorumCertificate, payloadHash flow.Identifier) flow.Identifier {
	id := make([]byte, 0)

	viewBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(viewBytes, view)
	id = append(id, viewBytes...)

	qcView := make([]byte, 8)
	binary.BigEndian.PutUint64(qcView, qc.View)
	id = append(id, qcView...)
	id = append(id, qc.BlockID[:]...)

	id = append(id, payloadHash[:]...)

	hasher := crypto.NewSHA3_256()
	hash := hasher.ComputeHash(id)

	var identifier flow.Identifier
	copy(identifier[:], hash)
	return identifier
}
