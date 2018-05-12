package gopool

// LocalDataFactory is a function template to allow a daemon to create a local data instance
// This instance exists for the lifetime of the daemon (not the gopool)
// This instance is unique to the daemon
type LocalDataFactory func() LocalData

// LocalData is an interface for setting up persistant data local to a daemon
type LocalData interface {
	// Setup is called to allow for any local data processes to occur
	// An error from this method will cause the daemon to exit
	Setup() error
}

// EOF
