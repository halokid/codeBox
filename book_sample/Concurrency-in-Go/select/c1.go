package main

import "fmt"

func main() {
  var c1, c2 <-chan interface{}
  var c3 chan<- interface{}

  select {
  case <- c1:
    fmt.Println("c1  writed")
  case <- c2:
    fmt.Println("c2 writed")
  case c3 <- struct{}{}:
    fmt.Println("c3 readed")
  }

}


