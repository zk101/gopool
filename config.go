package gopool

// Config holds static configuration items
//
// DaemonExpiry; The number of seconds a daemon will sit idle before exiting.
// MonitorLoopTimeout; The number of milliseconds of sleep between monitor goroutine cycles
// PoolSizeMin; The minimum number of daemons kept running
// PoolSizeMax; The maximum number of daemons in the pool
// QueueSize; The maximum number of Tasks allowed
// QueueTimeout; The number of milliseconds the Queue function will wait to queue a Task before returning a timeout
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
