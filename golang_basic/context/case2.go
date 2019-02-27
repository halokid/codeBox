package main

import (
  "fmt"
  "time"
)

func main() {
  stop := make(chan bool)

  go func() {
    for {
      select {
      case <- stop:
        fmt.Println("监控退出， 停止了.....")
        return
      default:
        fmt.Println("goroutine监控中")
        time.Sleep(2 * time.Second)
      }
    }
  }()

  time.Sleep(10 * time.Second)
  fmt.Println("可以了，通知监控停止")
  stop <- true
  time.Sleep(5 * time.Second)
}




