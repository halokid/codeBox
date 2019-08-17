package main

import (
  "fmt"
  "runtime"
  "sync"
)

func main() {
  memConsumed := func() uint64 {
    runtime.GC()
    var s runtime.MemStats
    runtime.ReadMemStats(&s)
    return s.Sys
  }

  var c <- chan interface{}   // 一个空的channel
  var wg sync.WaitGroup
  noop := func() { wg.Done(); <-c } // 1

  const numGoroutines = 1e4     // 2
  fmt.Println(numGoroutines)

  wg.Add(numGoroutines)
  before := memConsumed()  // 3

  for i := numGoroutines; i > 0; i-- {
    go noop()
  }
  wg.Wait()

  after := memConsumed()   // 4

  // 求新建立一个 goroutine 占多大内存， 因为上面建立了 numGoroutines 个 goroutine， 所以这里除以 numGoroutines
  fmt.Printf("%.3fkb", float64(after - before) / numGoroutines / 1000)
}







