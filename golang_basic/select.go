package main

import (
  "fmt"
)

func main() {
  ch := make(chan int, 3)
  ch <- 1
  
  // for {
  select {
    case ch <- 2:
    
    default:
      fmt.Println("channel is full")
  }
  // }
}