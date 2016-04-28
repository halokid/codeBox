package main 

import (
  "fmt"
)

var quit chan int = make(chan int)

func loop() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d ", i)
  }
  quit <- 0
}

func main() {
  go loop()
  go loop()
  
  for i := 0; i < 2; i++ {
    <- quit
  }
}