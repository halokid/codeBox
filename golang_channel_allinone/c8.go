package main

import (
  "fmt"
  "time"
)

/**
channel 超时处理
 */

func testTimeout() {
  g := make(chan int)
  quit := make(chan bool)

  go func() {
    for {
      select {
      case v := <-g:
        fmt.Println(v)
      case <-time.After(time.Second * time.Duration(5)):
        quit <-true
        fmt.Println("超时，通知主协程退出")
        return
      }
    }
  }()

  for i := 0; i < 3; i++ {
    g <-i
  }

  <-quit        // 如果quit没有写入，这里会一直阻塞
  fmt.Println("收到退出通知，主协程退出")
}

func main() {
  testTimeout()
}


