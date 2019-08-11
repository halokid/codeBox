package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

const maxWorkers = 5

type jobx struct {
  name        string
  duration    time.Duration
}

func doWorkx(id int, j jobx)  {
  fmt.Printf("worker%d: started %s, working for %fs\n", id, j.name, j.duration.Seconds())
  time.Sleep(j.duration)
  fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func main() {
  jobs := make(chan jobx)

  wg := &sync.WaitGroup{}
  //var wgx sync.WaitGroup

  // start workers, to do work

  // fixme: 假如 workers 是 5个， 那就是 5 个协程都去循环处理 jobs， 直到 jobs全部处理完为止
  // fixme: 值的注意的是，  jobs 是非缓冲channel， 那就是 写入一次， worker才能处理一次
  wg.Add(maxWorkers)
  for i := 1; i <= maxWorkers; i++ {
    go func(i int) {
      defer wg.Done()

			//fixme: 这里有一个关键的点就是， 这个for循环是去jobs这个channel拿数据，一直到消费完为止
      for j := range jobs {
        doWorkx(i, j)
      }
    }(i)
  }

  // add jobs
  iCk := 0
  for i := 0; i < 100; i++ {
    name := fmt.Sprintf("job-%d", i)
    duration := time.Duration(rand.Intn(1000)) * time.Millisecond
    fmt.Printf("adding: %s %s\n", name, duration)

    jobs <- jobx{name:  name, duration:  duration}     // add jobs

    iCk ++
  }
  close(jobs)

  // 等待 worker 完成处理
  wg.Wait()

  fmt.Println("golang会处理for循环写入channle的逻辑，因为这个是主进程，所以会确保主进程main的逻辑全部执行完 ------------- ", iCk)
}







