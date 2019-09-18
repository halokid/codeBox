package main

import (
  "fmt"
  "sync"
  "time"
)

func doWork(i int, done chan int) int {
  time.Sleep(6 * time.Second)
  fmt.Println("doWork ---------- ", i)
  if i == 6 {
    done <-i
  }
  return 1
}

func main() {
  wg := sync.WaitGroup{}
  done := make(chan int, 1)

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      doWork(i, done)
    }()
  }

  wg.Wait()

  for{
    select {
    case <-done:
      fmt.Println("done 1")
      //close(done)
      return
    }
  }

  fmt.Println("done 2...")
}

