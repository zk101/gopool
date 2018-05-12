package gopool

import "testing"

// TestService
func TestService(t *testing.T) {
	// Test nil config
	service, err := NewService(nil, nil)
	if err != nil {
		t.Fatalf("NewService failed: %s", err.Error())
	}

	service.Start()
	service.Stop()

	// Test pool size max < pool size min
	config := DefaultConfig()
	config.PoolSizeMax = 1
	config.PoolSizeMin = 10

	service, err = NewService(&config, nil)
	if err == nil {
		t.Fatalf("NewService succeeded where it should have failed")
	}

	// Test pool size less than 1
	config.PoolSizeMin = 0

	service, err = NewService(&config, nil)
	if err == nil {
		t.Fatalf("NewService succeeded where it should have failed")
	}
}

// EOF
