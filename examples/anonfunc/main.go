package main

import (
	"log"

	"github.com/zk101/gopool"
)

/*
  gopool.Task implementation
*/
type task struct {
	anonFunc  func()
	cancelled bool
	status    chan bool
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
	// Run our anonymous function
	t.anonFunc()
}

func (t *task) RunStatus(status bool) {
	t.status <- status
}

func main() {
	// Startup
	config := gopool.DefaultConfig()
	config.PoolSizeMax = 1

	service, err := gopool.NewService(&config, nil)
	if err != nil {
		log.Fatalf("New GoPool failed: %s\n", err.Error())
	}

	service.Start()

	// Run Anonymous Function
	message := "Hello World!"
	task := newTask()

	task.anonFunc = func() {
		log.Printf("%s\n", message)
	}

	if err := service.QueueTask(task); err != nil {
		log.Printf("Queuing a Task failed: %s\n", err.Error())
	}

	<-task.status

	// Shutdown
	service.Stop()
}

// EOF
