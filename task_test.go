package gopool

import (
	"time"
)

// A Task implementation for testing purposes

type testTask struct {
	cancelled bool
	status    chan bool
}

func newTestTask() *testTask {
	return &testTask{
		status: make(chan bool, 1),
	}
}

func (tt *testTask) IsCancelled() bool {
	return tt.cancelled
}

func (tt *testTask) Run(localData LocalData) {
	// Do some work
	time.Sleep(50 * time.Millisecond)
}

func (tt *testTask) RunStatus(status bool) {
	tt.status <- status
}

// EOF
