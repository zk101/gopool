package gopool

import (
	"errors"
	"time"
)

// QueueTask adds a Task to the pools queue
func (s *Service) QueueTask(task Task) error {
	if !s.run {
		return errors.New("gopool is not stated")
	}

	return s.queueTask(task, time.After(time.Millisecond*time.Duration(s.config.QueueTimeout)))
}

// QueueTaskWithTimeout adds a task with a custom timeout
func (s *Service) QueueTaskWithTimeout(task Task, timeout time.Duration) error {
	if !s.run {
		return errors.New("gopool is not stated")
	}

	return s.queueTask(task, time.After(timeout))
}

// queueTask adds a task to the queue
func (s *Service) queueTask(task Task, timeout <-chan time.Time) error {
	select {
	// This will occur if the queue is full and the register is full after the specified timeout
	case <-timeout:
		return errors.New("timeout queueing task")

	// Queue the task
	case s.queue <- task:
		return nil

	// Scale the pool
	case s.register <- struct{}{}:
		go s.daemon(task)
		return nil
	}
}

// EOF
