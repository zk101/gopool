package gopool

// LocalDataFactory is a function template to allow a daemon to create a local data instance
// This instance exists for the lifetime of the daemon (not the gopool)
// This instance is unique to the daemon
type LocalDataFactory func() LocalData

// LocalData is an interface for setting up persistant data local to a daemon
type LocalData interface {
	// Start is called to setup and start up any data required for processing tasks
	// An error from this method will cause the daemon to exit
	Start() error

	// Stop is called when a daemon terminates cleanly to close down any services and perform any cleanup
	// Stop does not return an error as doing so is rather redundant.  The implementation should deal with that
	Stop()
}

// EOF
