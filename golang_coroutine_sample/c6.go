package main

import (
  "fmt"
  "math/rand"
  "time"
)

func send(ch chan int) {
  for {
    var value = rand.Intn(100)
    ch <- value
    fmt.Printf("send %d\n", value)
  }
}

func recv(ch chan int) {
  for {
    value := <- ch
    fmt.Printf("recv %d\n", value)
    time.Sleep(time.Second)
  }
}

func main() {
  var ch = make(chan int, 1)

  go recv(ch)

  send(ch)
}





