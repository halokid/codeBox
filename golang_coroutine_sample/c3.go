package main

import (
  "fmt"
  "runtime"
  "time"
)

func main() {
  fmt.Println(runtime.NumGoroutine())

  for i := 0; i < 10; i++ {
    go func() {
      for {
        time.Sleep(time.Second)
      }
    }()
  }

  // 产生的 10 个 grontine， 加上main进程一个， 一共 11 个
  fmt.Println(runtime.NumGoroutine())
}
