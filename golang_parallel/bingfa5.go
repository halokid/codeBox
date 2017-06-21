package main

import (
  "fmt"
  "time"
)

func A(c chan int) {
  for i := 0; i < 10; i++ {
    c <- i
  }
}

func B(c chan int) {
  for val := range c {
    fmt.Println("value: ", val)
  }
}


func main() {
  chs := make(chan int, 10)

  go A(chs)
  go B(chs)

  time.Sleep(1e9)
}
