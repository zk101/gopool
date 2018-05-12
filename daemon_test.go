package gopool

import (
	"testing"
	"time"
)

// TestDaemonTimout checks the daemon times out after the expiry period
func TestDaemonTimeout(t *testing.T) {
	config := DefaultConfig()

	// Set expiry to one second
	config.DaemonExpiry = 1

	service, err := NewService(&config, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.run = true
	daemonExited := make(chan struct{}, 1)

	go func() {
		// Pump the register - Not doing this will cause the daemon to block on exit
		service.register <- struct{}{}
		service.daemon(nil)
		daemonExited <- struct{}{}
	}()

	// Allow the goroutine to start up
	time.Sleep(5 * time.Millisecond)

	select {
	// If we are still running after two seconds, daemon failed to exit
	case <-time.After(2 * time.Second):
		service.run = false
		service.wg.Wait()
		t.Error("Daemon failed to timeout")

	case <-daemonExited:
		// Daemon exiting before timeout is good
	}
}

// TestDaemonRunFalse checks the daemon exits immediately if run is false
func TestDaemonRunFalse(t *testing.T) {
	service, err := NewService(nil, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	daemonExited := make(chan struct{}, 1)

	go func() {
		// Pump the register - Not doing this will cause the daemon to block on exit
		service.register <- struct{}{}
		service.daemon(nil)
		daemonExited <- struct{}{}
	}()

	// Allow the goroutine to start up
	time.Sleep(5 * time.Millisecond)

	select {
	case <-time.After(50 * time.Millisecond):
		service.run = false
		service.wg.Wait()
		t.Error("Daemon failed to exit with run set to false")

	case <-daemonExited:
		// Worker exiting before timeout is good
	}
}

// TestDaemonRunTask works
func TestDaemonRunTask(t *testing.T) {
	service, err := NewService(nil, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.run = true

	// Test daemon runs a task provided
	task := newTestTask()

	go func() {
		service.register <- struct{}{}
		service.daemon(task)
	}()

	if status := <-task.status; !status {
		t.Error("Task failed to run")
	}

	service.run = false
	service.wg.Wait()
}

// TestDaemonCancelledTask works
func TestDaemonCancelledTask(t *testing.T) {
	service, err := NewService(nil, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.run = true

	// Test daemon runs a task provided
	task := newTestTask()
	task.cancelled = true

	go func() {
		service.register <- struct{}{}
		service.daemon(task)
	}()

	if status := <-task.status; status {
		t.Error("Task run but was cancelled")
	}

	service.run = false
	service.wg.Wait()
}

// TestDaemonLocalData
func TestDaemonLocalData(t *testing.T) {
	service, err := NewService(nil, testLocalDataFactory)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.run = true
	task := newTestTask()

	go func() {
		service.register <- struct{}{}
		service.daemon(task)
	}()

	if status := <-task.status; !status {
		t.Error("Task failed to run")
	}

	service.run = false
	service.wg.Wait()
}

// TestDaemonLocalDataWithError
func TestDaemonLocalDataWithError(t *testing.T) {
	service, err := NewService(nil, testLocalDataFactoryWithError)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.run = true
	task := newTestTask()

	go func() {
		service.register <- struct{}{}
		service.daemon(task)
	}()

	if status := <-task.status; status {
		t.Error("Task ran but localdata had an error")
	}

	service.run = false
	service.wg.Wait()
}

// EOF
