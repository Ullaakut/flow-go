package dps

// Config represents the configuration of the DPS servers.
type Config struct {
	Host     string
	PubPort  uint16
	SyncPort uint16
}
