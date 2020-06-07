package main

import (
  "fmt"
  "time"
)

/**
线程抢占输出
 */

func testSequnse()  {
  ch := make(chan int)

  go func() {
    v := <-ch
    fmt.Println("协程1输出v:", v)
  }()

  ch <-1
  fmt.Println("协程0输出2")
  time.Sleep(time.Second)
}

func main() {
  testSequnse()
}
