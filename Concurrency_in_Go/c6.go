package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func AddVal(i int) {
  i++
  fmt.Println(i)
  wg.Done()
}

func main() {
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go AddVal(i)
  }
  wg.Wait()
}
