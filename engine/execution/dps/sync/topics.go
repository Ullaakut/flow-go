package sync

// Topics are used as headers to all messages that transit over the sync socket.
// They allow the server to know how to parse synchronization requests.
var (
	TopicBlock = []byte(`sync_block`)
	TopicTrie = []byte(`sync_trie`)
)
