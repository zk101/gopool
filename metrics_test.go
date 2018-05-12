package gopool

import "testing"

// TestMetrics
func TestMetrics(t *testing.T) {
	service, err := NewService(nil, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.Start()

	// Check using default config, we should have a single daemon running
	if service.GetDaemonCount() != 1 {
		t.Errorf("Found daemon count of %d, expected %d", service.GetDaemonCount(), 1)
	}

	// Check we have no tasks in the queue
	if service.GetQueueCount() != 0 {
		t.Errorf("Found queue count of %d, expected %d", service.GetQueueCount(), 0)
	}

	service.Stop()
}

// EOF
