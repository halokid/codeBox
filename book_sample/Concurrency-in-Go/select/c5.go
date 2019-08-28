package main

import (
  "fmt"
  "time"
)

var suss chan int
//suss := make(chan int)
var fail chan int

func doWork()  {
  for i := 0; i < 10; i++ {
    suss <-i
    fmt.Println(i)
  }

  // fail
  fail <-1

  for j := 0; j < 10; j++ {
    suss <-j
  }
}

func main() {

  go doWork()

  for {
    select {
    case <-suss:
      fmt.Println("xx")

    }
  }

  time.Sleep(5 * time.Second)
  fmt.Println("break for loop, done")

}





