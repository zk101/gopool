package gopool

// GetDaemonCount returns the current number of running daemons routines
func (s *Service) GetDaemonCount() int {
	return len(s.register)
}

// GetQueueCount returns the current number of tasks on the queue
func (s *Service) GetQueueCount() int {
	return len(s.queue)
}

// EOF
