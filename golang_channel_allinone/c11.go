package main

import (
  "fmt"
  "os"
  "os/signal"
  "time"
)

/** 监听中断信号的channel */

func testSignal() {
  quit := make(chan os.Signal)
  signal.Notify(quit, os.Interrupt)

  go func() {
    time.Sleep(time.Second * 2)

    number := 0
    for {
      number++
      fmt.Println("number:", number)
      time.Sleep(time.Second)
    }
  }()

  fmt.Println("按Ctrl+C可退出程序")
  <-quit
  fmt.Println("协程退出")
}

func main() {
  testSignal()
}
