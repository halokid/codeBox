package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

func main() {
  doWork := func(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan <-int ) {
    started := time.Now()
    defer wg.Done()

    simulatedLoadTime := time.Duration(1 + rand.Intn(5)) * time.Second

    select {
    case <-done:
    case <-time.After(simulatedLoadTime):
    }

    select {
    case <-done:
    case result <-id:
    }

    took := time.Since(started)
    if took < simulatedLoadTime {
      took = simulatedLoadTime
    }
    fmt.Printf("%v took %v\n", id, took)
  }

  done := make(chan interface{})
  result := make(chan int)

  var wg sync.WaitGroup
  wg.Add(10)
  for i := 0; i < 10; i++ {
    go doWork(done, i, &wg, result)
  }

  firstReturn := <-result

  i := 1
  for res := range result {
    i++
    fmt.Println("并行处理获取结果:  ", res)
    if i == 10 {
      break
    }
  }

  close(done)
  wg.Wait()
  fmt.Printf("Received an answer from #%v\n", firstReturn)
}



