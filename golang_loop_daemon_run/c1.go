package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("hello, world")
  a := 1e8
  fmt.Printf("%f\n", a)
  loopWorker()
}

func loopWorker() {
  i := 0
  ticker := time.NewTicker(6 * time.Second)
  defer ticker.Stop()

  for {
    select {
    case <- ticker.C:
      i++
      doxx(i)
    }
  }
}

func doxx(i int) {
  //time.Sleep(7 * time.Second)
  fmt.Println("doxx ", i, time.Now().Unix())
}


