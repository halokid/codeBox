package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan int, 1)

  for {
    select {
      case ch <- 0:
      case ch <- 1:
    }

    i := <- ch
    fmt.Println("value receive: ", i)
    time.Sleep(10 * time.Second)
  }
}
