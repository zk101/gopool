package gopool

// Task interface
type Task interface {
	// IsCancelled should return true if the task should no longer be processed, otherwise return false
	IsCancelled() bool

	// Run is called to execute the Task.  It passes a localdata pointer
	// It is up the the implementation to handle a nil pointer
	Run(LocalData)

	// RunStatus is used to notify when a task ihas finished being run
	// false indicates the daemon failed to run the task (an error occured in LocalData setup)
	// true indicates the task ran (runTask completed)
	RunStatus(bool)
}

// EOF
