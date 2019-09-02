package main

import (
"fmt"
"sync"
"time"
)

func wgDoWork(wg *sync.WaitGroup) {
  defer wg.Done()
  time.Sleep(5 * time.Millisecond)
  fmt.Println("wgDoWork...")
}

func main() {
  var wg sync.WaitGroup

  start := time.Now()
  for i := 0; i < 100000; i++ {
    wg.Add(1)
    go wgDoWork(&wg)
  }

  wg.Wait()
  fmt.Println("time cost:  ", time.Since(start))
}


