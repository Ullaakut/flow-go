package pub

import (
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/model/flow"
)

// Notification represents a change of state that needs to be published to the
// subscribed DPS instances in order to keep them up-to-date with the network.
// Notifications can contain new blocks or trie updates.
type Notification struct {
	typ notificationType

	// Trie updates.
	view     state.View
	rootHash ledger.RootHash

	// Blocks.
	block *flow.Block

	// TODO: Handle events (https://github.com/optakt/flow-dps/issues/99)
}

type notificationType uint8

const (
	notificationTypeTrieUpdate notificationType = 1
	notificationTypeBlock      notificationType = 2
)

// NotifyTrieUpdate sends an appropriately-typed notification to a given notification channel,
// containing the given state view and root hash.
func NotifyTrieUpdate(notifyCh chan<- Notification, view state.View, rootHash ledger.RootHash) {
	notifyCh <- Notification{
		typ: notificationTypeTrieUpdate,

		view:     view,
		rootHash: rootHash,
	}
}

// NotifyBlock sends an appropriately-typed notification to a given notification channel,
// containing the given block pointer.
func NotifyBlock(notifyCh chan<- Notification, block *flow.Block) {
	notifyCh <- Notification{
		typ: notificationTypeBlock,

		block: block,
	}
}
