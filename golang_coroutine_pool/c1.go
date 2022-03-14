package main

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





