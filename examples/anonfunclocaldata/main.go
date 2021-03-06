package main

import (
	"errors"
	"log"

	"github.com/zk101/gopool"
)

/*
  gopool.LocalData implementation
*/
func exampleLocalDataFactory() gopool.LocalData {
	return &exampleLocalData{}
}

type exampleLocalData struct {
	count int
}

func (eld *exampleLocalData) Start() error {
	eld.count = 10

	return nil
}

func (eld *exampleLocalData) Stop() {}

/*
  gopool.Task implementation
*/
type task struct {
	anonFunc  func(*exampleLocalData) error
	cancelled bool
	status    chan bool
	err       error
}

func newTask() *task {
	return &task{
		status: make(chan bool, 1),
	}
}

func (t *task) IsCancelled() bool {
	return t.cancelled
}

func (t *task) Run(localData gopool.LocalData) {
	if localData == nil {
		t.err = errors.New("localdata is nil")
		return
	}

	if err := t.anonFunc(localData.(*exampleLocalData)); err != nil {
		t.err = err
		return
	}
}

func (t *task) RunStatus(status bool) {
	t.status <- status
}

func main() {
	// Startup
	config := gopool.DefaultConfig()
	config.PoolSizeMax = 1

	service, err := gopool.NewService(&config, exampleLocalDataFactory)
	if err != nil {
		log.Fatalf("New GoPool failed: %s\n", err.Error())
	}

	service.Start()

	for x := 0; x < 10; x++ {
		// Run a task with localdata
		task := newTask()
		task.anonFunc = func(data *exampleLocalData) error {
			log.Printf("%d\n", data.count)
			data.count++
			return nil
		}

		if err := service.QueueTask(task); err != nil {
			log.Printf("Queuing a Task failed: %s\n", err.Error())
		}

		<-task.status

		if task.err != nil {
			log.Printf("Found task error: %s\n", task.err.Error())
		}
	}

	// Shutdown
	service.Stop()
}

// EOF
