package gopool

import "testing"

// TestDefaultConfig checks we are running sane values
func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.DaemonExpiry != 30 {
		t.Errorf("Expected for config.Expiry value %d but got value %d", 30, config.DaemonExpiry)
	}

	if config.MonitorLoopTimeout != 50 {
		t.Errorf("Expected for config.MonitorLoopTimeout value %d but got value %d", 50, config.MonitorLoopTimeout)
	}

	if config.PoolSizeMin != 1 {
		t.Errorf("Expected for config.Min value %d but got value %d", 1, config.PoolSizeMin)
	}

	if config.PoolSizeMax != 10 {
		t.Errorf("Expected for config.Max value %d but got value %d", 10, config.PoolSizeMax)
	}

	if config.QueueSize != 10 {
		t.Errorf("Expected for config.QueueSize value %d but got value %d", 10, config.QueueSize)
	}

	if config.QueueTimeout != 50 {
		t.Errorf("Expected for config.QueueTimeout value %d but got value %d", 50, config.QueueTimeout)
	}
}

// EOF
