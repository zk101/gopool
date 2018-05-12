package gopool

// Config holds static configuration items
type Config struct {
	DaemonExpiry       int
	MonitorLoopTimeout int
	PoolSizeMin        int
	PoolSizeMax        int
	QueueSize          int
	QueueTimeout       int
}

// DefaultConfig returns a Config struct with predefined settings
func DefaultConfig() Config {
	return Config{
		DaemonExpiry:       30,
		MonitorLoopTimeout: 50,
		PoolSizeMin:        1,
		PoolSizeMax:        10,
		QueueSize:          10,
		QueueTimeout:       50,
	}
}

// EOF
