package gopool

import (
	"time"
)

// daemon is the task processor goroutine
func (s *Service) daemon(task Task) {
	s.wg.Add(1)

	var localData LocalData

	defer func() {
		recover()

		if localData != nil {
			localData.Stop()
		}

		<-s.register
		s.wg.Done()
	}()

	if !s.run {
		return
	}

	lastTask := time.Now().Add(time.Second * time.Duration(s.config.DaemonExpiry))
	monitorChan := make(chan bool, 1)
	loopControl := true

	// Create and start LocalData if applicable
	if s.localDataFactory != nil {
		localData = s.localDataFactory()
	}

	if localData != nil {
		if err := localData.Start(); err != nil {
			if task != nil {
				task.RunStatus(false)
			}

			return
		}
	}

	// Start daemon monitor routine
	go func() {
		for {
			if time.Now().After(lastTask) {
				monitorChan <- true
				break
			}

			if !s.run {
				monitorChan <- true
				break
			}

			// Prevent our monitor routine trashing the cpu
			time.Sleep(time.Millisecond * time.Duration(s.config.MonitorLoopTimeout))
		}
	}()

	if task != nil {
		s.runTask(task, localData)
	}

	// Daemon Run Loop
	for loopControl {
		select {
		case <-monitorChan:
			loopControl = false

		case task := <-s.queue:
			s.runTask(task, localData)
			lastTask = time.Now().Add(time.Second * time.Duration(s.config.DaemonExpiry))
		}
	}
}

func (s *Service) runTask(task Task, localData LocalData) {
	if task.IsCancelled() {
		task.RunStatus(false)
		return
	}

	task.Run(localData)

	task.RunStatus(true)
}

// EOF
