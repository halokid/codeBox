package main

import (
  "fmt"
)

func main() {
  var ch = make(chan int, 4)
  ch <- 1
  ch <- 2
  close(ch)

  for value := range ch {
    fmt.Println(value)
  }
}
