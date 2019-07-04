package main

import (
  "flag"
  "fmt"
  "log"
  "net/http"
  "time"
)

type Job struct {
  Name        string
  Delay       time.Duration
}

type Worker struct {
  id          int
  jobQueue    chan Job
  workerPool  chan chan Job
  quitChan    chan bool
}

func NewWorker(id int, workerPool chan chan Job) Worker {
  return Worker{
    id:                 id,
    jobQueue:           make(chan Job),
    workerPool:         workerPool,
    quitChan:           make(chan bool),
  }
}

func (w Worker) start() {
  go func() {
    for {
      w.workerPool <- w.jobQueue        // 只是把 jobQueue  的值 导向到另外一个 channel， 是复制的意思, 等于是  workerPool 的channel 写入了一个值， 事实上 jobQueue 还是需要一个值去读取它的

      select {
      case job := <- w.jobQueue:
        fmt.Printf("worker %d: started %s, blocking for %f second\n", w.id, job.Name, job.Delay.Seconds())
        time.Sleep(job.Delay)
        fmt.Printf("worker %d: completed %s!\n", w.id, job.Name)
      case <- w.quitChan :
        fmt.Printf("worker %d stopping\n", w.id)
        return
      }
    }
  }()
}

func (w Worker) stop() {
  go func() {
    w.quitChan <- true
  }()
}


type Dispatcher struct {
  workerPool          chan chan Job
  maxWorkers           int
  jobQueue            chan Job
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
  workerPool := make(chan chan Job, maxWorkers)

  return &Dispatcher{
    jobQueue:           jobQueue,
    maxWorkers:         maxWorkers,
    workerPool:         workerPool,
  }
}

func (d *Dispatcher) run() {
  for i := 0; i < d.maxWorkers; i++ {
    worker := NewWorker(i + 1, d.workerPool)
    worker.start()
  }

  d.dispatch()
}

func (d *Dispatcher) dispatch() {
  for {
    select {
    case job := <-d.jobQueue:
      go func() {
        fmt.Printf("fetching workerJobQueue for: %s\n", job.Name)
        workerJobQueue := <-d.workerPool
        fmt.Printf("adding %s to workerJobQueue\n", job.Name)
        workerJobQueue <- job
      }()
    }
  }
}


func requestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
  // Make sure we can only be called with an HTTP POST request.
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  // Parse the delay.
  delay, err := time.ParseDuration(r.FormValue("delay"))
  if err != nil {
    http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
    return
  }

  // Validate delay is in range 1 to 10 seconds.
  if delay.Seconds() < 1 || delay.Seconds() > 10 {
    http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
    return
  }

  // Set name and validate value.
  name := r.FormValue("name")
  if name == "" {
    http.Error(w, "You must specify a name.", http.StatusBadRequest)
    return
  }

  // Create Job and push the work onto the jobQueue.
  job := Job{Name: name, Delay: delay}
  jobQueue <- job

  // Render success.
  w.WriteHeader(http.StatusCreated)
}

func main() {
  var (
    maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
    maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
    port         = flag.String("port", "8080", "The server port")
  )
  flag.Parse()

  // Create the job queue.
  jobQueue := make(chan Job, *maxQueueSize)

  // Start the dispatcher.
  dispatcher := NewDispatcher(jobQueue, *maxWorkers)
  dispatcher.run()

  // Start the HTTP handler.
  http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
    requestHandler(w, r, jobQueue)
  })
  log.Fatal(http.ListenAndServe(":"+*port, nil))
}





















