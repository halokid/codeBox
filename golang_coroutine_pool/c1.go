package main

import (
  "fmt"
  "time"
)

type Task struct {
  f func() error
}

func NewTask(f func() error) *Task {
  t := Task{
    f: f,
  }
  return &t
}

func (t *Task) Execute() {
  t.f()
}

type Pool struct {
  EntryChannel chan *Task

  workerNum int

  JobsChannel chan *Task
}

func NewPool(cap int) *Pool {
  p := Pool{
    EntryChannel:  make(chan *Task),
    workerNum:  cap,
    JobsChannel:  make(chan *Task),
  }

  return &p
}

// the pool create a worker and start work
func (p *Pool) worker(workerID int) {
  for task := range p.JobsChannel {
    task.Execute()
    fmt.Println("workerID ", workerID, " process done.")
  }
}

// make the pool start work
func (p *Pool) Run() {
  // todo: 1. make the workder for process job
  for i := 0; i < p.workerNum; i++ {
    go p.worker(i)
  }

  // todo: 2. get task from EntryChannel transmit by outside
  // todo: and send the job into JobsChannel
  for task := range p.EntryChannel {
    p.JobsChannel <- task
  }

  // todo: 3. process done to close JobsChannel
  close(p.JobsChannel)

  // todo: 4. process done to close EntryChannle
  close(p.EntryChannel)
}

func main() {
  // create a task
  task := NewTask(func() error {
    fmt.Println(time.Now())
    return nil
  })

  // create a pool, set 3 worker
  pool := NewPool(3)

  // create a coroutine, constanly send task to Pool
  go func() {
    for {
      pool.EntryChannel <- task
      time.Sleep(2 * time.Second)
    }
  }()

  // start the pool to work
  pool.Run()
}





