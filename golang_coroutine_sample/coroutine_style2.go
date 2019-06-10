package main

import (
  "fmt"
  "runtime"
  "time"
)

/**
显式声明让出cpu， 两个 goroutine 不会抢占 CPU
 */

var quitx = make(chan int)

func loopx() {
  for i := 0; i < 10; i++ {
    runtime.Gosched()
    fmt.Printf("%d\n", i)
  }

  quitx <- 0
}

func main() {
  t1 := time.Now()
  runtime.GOMAXPROCS(2)

  go loopx()
  go loopx()

  for i := 0; i < 2; i++ {
    <- quitx
  }

  t2 := time.Now()
  fmt.Println(t2.Sub(t1))
}


