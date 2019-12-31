package main

import (
  "fmt"
  "time"
)

func main() {
  requests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests <-i
  }
  close(requests)

  limiter := time.Tick(200 * time.Millisecond)

  // 1 秒跑 5 个
  for req := range requests {
    <-limiter        // 200 millisecond 写入一次channel， 否则阻塞
    fmt.Println("接收到请求", req, time.Now())
  }

  // ----------------------------------------------------------------------------------------

  burstyLimiter := make(chan time.Time, 3)
  for i := 0; i < 3; i++ {
    burstyLimiter <-time.Now()
  }

  go func() {
    for t := range time.Tick(200 * time.Millisecond) {
      burstyLimiter <-t
    }
  }()

  burstyRequests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    burstyRequests <-i
  }
  close(burstyRequests)

  for req := range burstyRequests {
    <-burstyLimiter
    fmt.Println("request", req, time.Now())
  }

}
