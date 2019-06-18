package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

const maxWorkers = 5

type job struct {
  name        string
  duration    time.Duration
}

func doWork(id int, j job)  {
  fmt.Printf("worker%d: started %s, working for %fs\n", id, j.name, j.duration.Seconds())
  time.Sleep(j.duration)
  fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func main() {
  jobs := make(chan job)

  wg := &sync.WaitGroup{}
  //var wgx sync.WaitGroup

  // start workers, to do work
  wg.Add(maxWorkers)
  for i := 1; i <= maxWorkers; i++ {
    go func(i int) {
      defer wg.Done()

      for j := range jobs {
        doWork(i, j)
      }
    }(i)
  }

  // add jobs
  for i := 0; i < 100; i++ {
    name := fmt.Sprintf("job-%d", i)
    duration := time.Duration(rand.Intn(1000)) * time.Millisecond
    fmt.Printf("adding: %s %s\n", name, duration)

    jobs <- job{name:  name, duration:  duration}     // add jobs
  }
  close(jobs)

  // 等待 worker 完成处理
  wg.Wait()
}







