package gopool

import "testing"

// TestQueue
// queueTask select will randomly choose between starting a daemon and using the queue, so test coverage is not always 100%
func TestQueue(t *testing.T) {
	service, err := NewService(nil, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	// Tasks should not queue if the pool isn't running
	if err := service.QueueTask(newTestTask()); err == nil {
		t.Error("Task was queued, however the gopool is not running")
	}

	if err := service.QueueTaskWithTimeout(newTestTask(), 5); err == nil {
		t.Error("Task was queued, however the gopool is not running")
	}

	service.Start()

	// Tasks should queue if the gopool is running
	if err := service.QueueTask(newTestTask()); err != nil {
		t.Errorf("Queue task failed: %s", err.Error())
	}

	// Test queuing with timeout
	if err := service.QueueTaskWithTimeout(newTestTask(), 5); err != nil {
		t.Errorf("Queue task failed: %s", err.Error())
	}

	service.Stop()

	// Test Timeout works, start with a minimising the pool
	config := DefaultConfig()
	config.PoolSizeMax = 1
	config.QueueSize = 0

	service, err = NewService(&config, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.Start()

	// Queue a task (Run sleeps for 50ms)
	if err := service.QueueTask(newTestTask()); err != nil {
		t.Errorf("Queue task failed: %s", err.Error())
	}

	// Queue another task (our actual test)
	if err := service.QueueTaskWithTimeout(newTestTask(), 5); err == nil {
		t.Error("Task should have timed out")
	}

	service.Stop()
}

// EOF
