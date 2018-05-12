package gopool

import (
	"errors"
	"sync"
	"time"
)

// Service holds operational configuration items and methods
type Service struct {
	config           *Config
	localDataFactory LocalDataFactory
	queue            chan Task
	register         chan struct{}
	run              bool
	wg               *sync.WaitGroup
}

// NewService sets up a new Service struct
func NewService(conf *Config, localDataFactory LocalDataFactory) (*Service, error) {
	if conf == nil {
		c := DefaultConfig()
		conf = &c
	}

	if conf.PoolSizeMin < 1 {
		return nil, errors.New("minimum pool size must be one or greater")
	}

	if conf.PoolSizeMin > conf.PoolSizeMax {
		return nil, errors.New("minimum pool size must be less than or equal to the maximum pool size")
	}

	return &Service{
		config:           conf,
		localDataFactory: localDataFactory,
		queue:            make(chan Task, conf.QueueSize),
		register:         make(chan struct{}, conf.PoolSizeMax),
		wg:               &sync.WaitGroup{},
	}, nil
}

// Start launches the pool monitor
func (s *Service) Start() {
	// Set run to true, this is used to control all the goroutines
	s.run = true

	// Start the monitor routine
	go s.monitor()

	// Allow the monitor to get the pool running
	time.Sleep(time.Millisecond * 5)
}

// Stop sets run to false and waits for all the daemons to halt
func (s *Service) Stop() {
	s.run = false
	s.wg.Wait()
}

// monitor controls the daemon pool, ensuring minimum pool size is always running
func (s *Service) monitor() {
	for s.run == true {
		for x := len(s.register); x < s.config.PoolSizeMin; x++ {
			s.register <- struct{}{}
			go s.daemon(nil)
		}

		// Prevent monitor from thrashing the cpu
		time.Sleep(time.Millisecond * time.Duration(s.config.MonitorLoopTimeout))
	}
}

// EOF
