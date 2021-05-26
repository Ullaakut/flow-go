package pub

// Topics are used as headers to all messages that transit over the publish socket.
// They allow consumers to filter out unwanted messages.
var (
	TopicBlock = []byte(`pub_block`)
	TopicTrie = []byte(`pub_trie`)
)
