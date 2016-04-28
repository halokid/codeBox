package main 

import (
  "fmt"
  "time"
)

var quit chan int

func foo(id int) {
  fmt.Println(id)
  time.Sleep(time.Second)
  quit <- 0 
}

func main() {
  count := 1000
  quit = make(chan int, count)
  
  for i := 0; i < count; i++ {
    go foo(i)
  }
  
  for i :=0; i < count; i++ {
    <- quit
  }
}

