package main

import (
  "fmt"
  "time"
)

func testMaxNumControl() {
  maxNum := 3
  limit := make(chan bool, maxNum)      // 并发量
  quit := make(chan bool)

  for i := 0; i < 100; i++ {
    fmt.Println("开始------------:", i)

    limit <-true        // limit进3个则满，开始阻塞limit，直到有limit出栈， <-limit才放行

    go func(i int) {
      <-limit       // go func是并行处理的，因为 <-limit 的阻塞， 所以limit出栈一个则go func执行一个

      fmt.Println("执行------------:", i)
      time.Sleep(time.Millisecond * 20)
      fmt.Println("完成------------:", i)

      if i == 99 {
        fmt.Println("完成任务")
        quit <-true
      }
    }(i)
  }

  <-quit
  fmt.Println("收到退出通知， 主协程退出")
}

func main() {
  testMaxNumControl()
}


