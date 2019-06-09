package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

var (
  ops    uint64
  wg     sync.WaitGroup
)

func counter()  {
  defer wg.Done()
  for i := 0; i < 2; i++ {
    atomic.AddUint64(&ops, 1)
  }
}

func main() {
  wg.Add(2)
  go counter()
  go counter()
  wg.Wait()
  fmt.Println(ops)
}