package main
/**
run app:
go run worker_original.go -max_workers 5

request test:
for i in {1..15}; do curl localhost:8080/work -d name=job$i -d delay=$(expr $i % 9 + 1)s; done

 */
import (
  "configcenter/src/framework/core/log"
  "flag"
  "fmt"
  "net/http"
  "time"
)

type job struct {
  name        string
  duration    time.Duration
}


func doWork(id int, j job) {
  // worker真正处理的逻辑
  fmt.Printf("worker %d: started %s, working for %f seconds\n", id, j.name, j.duration.Seconds())
  time.Sleep(j.duration)
  fmt.Printf("worker %d: completed %s!\n", id, j.name)
}


func requestHandler(jobs chan job, w http.ResponseWriter, r *http.Request) {
  // 添加 job
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  duration, err := time.ParseDuration(r.FormValue("delay"))
  if err != nil {
    http.Error(w, "Bad deplay value: " + err.Error(), http.StatusBadRequest)
    return
  }

  if duration.Seconds() < 1 || duration.Seconds() > 10 {
    http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
    return
  }

  name := r.FormValue("name")
  if name == "" {
    http.Error(w, "You must specify a name.", http.StatusBadRequest)
    return
  }

  // 接收 http 请求， 把job写入jobs
  job := job{name: name, duration:  duration}
  go func() {
    fmt.Printf("added: %s %s\n", job.name, job.duration)
    jobs <- job
  }()

  w.WriteHeader(http.StatusCreated)
  return
}

func main() {
  var (
    maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
    maxWorkders = flag.Int("max_workers", 5, "The number of workers to start")
    port = flag.String("port", "8080", "The server port")
  )

  flag.Parse()

  jobs := make(chan job, *maxQueueSize)

  // fixme: 循环 workers， 通过协程把 job 添加到 doWork里面去处理, 所以这几个 worker 是同时处理 job 的
  // fixme: jobs channle 的写入由另外一个逻辑去负责
  for i := 1; i <= *maxWorkders; i++ {
    go func(i int) {
      // 这里会阻塞， 直到有新的job读取，才会执行，开一个新的 gor 来执行这个逻辑，貌似比select要清晰一点？？
      for j := range jobs {
        doWork(i, j)
      }
    }(i)
  }

  http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
    requestHandler(jobs, w, r)
  })
  log.Fatal(http.ListenAndServe(":" + *port, nil))
}




