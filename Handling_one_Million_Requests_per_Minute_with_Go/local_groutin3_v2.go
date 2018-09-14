
var (
  max_workers = os.Getenv("MAX_WORKERS")
  max_queue = os.Getenv("MAX_QUEUE")
)



// 工作请求结构体
type Job struct {
  payload  Payload 
}


// 工作请求
var job_queue chan Job


// 处理工作 worker 的结构体
type Worker struct {
  worker_pool chan chan Job
  job_channel chan Job
  quit    chan bool
}



func NewWorker(worker_pool chan chan Job) Worker {
  return Worker {
    worker_pool:    worker_pool,
    job_channel:    make(chan Job),
    quit:           make(chan bool),
  }
}


func (w Worker) Start() {
  go func() {
    // 守护监听 channel 的IO信息
    for {
      w.worker_pool <- w.job_channel

      // 监听判断 worker 是写入job， 还是写入quit
      select {
      // 取出 job
      case job := <- w.job_channel:
        if err := job.Payload.UploadToS3(); err != nil {
          log.Error("error uploading to S3: %s", err.Error())
        }

      case <- w.quit:
        return

      }
    }
  }
}


func (w Worker) Stop() {
  go func() {
    w.quit <- true
  }
}





func payloadHandler(w http.ResponseWriter, r *http.Request) {
  // golang 都是习惯先进行错误判断的。。。
  if r.Method != "POST" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  var content = &PayloadCollection{}
  err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
  if err != nil {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  for _, payload := range content.Payloads {
    work := Job{Payload: payload}

    job_queue <-  work
  }

  w.WriteHeader(http.StatusOK)

}




func main() {
  dispathcher := NewDispatcher(MaxWorker)
  dispathcher.Run()
}



type Dispathcher struct {
  worker_pool chan chan Job
}


func NewDispatcher(max_workers int) *Dispathcher {
  pool := make(chan chan Job, max_workers)
  return &Dispathcher{worker_pool:  pool}
}


func (d *Dispathcher) Run() {
  for i := 0; i < max_workers; i++ {
    worker := NewWorker(d.pool)
    worker.Start()
  }

  go d.dispathch()
}




































