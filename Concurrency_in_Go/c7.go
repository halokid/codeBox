package main

import (
  "fmt"
)

var c chan int

func AddValx(i int) {
  i++
  fmt.Println(i)
  c <- i
}


func main() {
  c = make(chan int)

  for i := 0; i < 10; i++ {
    go AddValx(i)
  }

  select {
  case cx := <- c:
    fmt.Println("cx ------- ", cx)
  default:
    fmt.Println("default ------- ")

  }
}

