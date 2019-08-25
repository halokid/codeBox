package main

import (
  "fmt"
  "time"
)

func main() {
  var c <-chan int      // 读取channel， 只能读取
  select {
  case <-c:
  case <-time.After(1 * time.Second):
    fmt.Println("channel time out")

  }
}
