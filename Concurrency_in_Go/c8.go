package main

import (
  "fmt"
  "sync"
)

var count int
var lock sync.Mutex

func main() {

  increment := func() {
    lock.Lock()
    defer lock.Unlock()
    count++
    fmt.Printf("Increment:   %d\n", count)
  }

  decrement := func() {
    lock.Lock()
    defer lock.Unlock()
    count--
    fmt.Printf("Decrement:     %d\n", count)
  }

  // increment
  var arithmetic sync.WaitGroup
  for i := 0; i <= 5; i++ {
    arithmetic.Add(1)
    go func() {
      defer arithmetic.Done()
      increment()
    }()
  }

  // decrement
  for i := 0; i <= 5; i++ {
    arithmetic.Add(1)
    go func() {
      defer arithmetic.Done()
      decrement()
    }()
  }

  arithmetic.Wait()
  fmt.Println(" ------ finished ------")
}




