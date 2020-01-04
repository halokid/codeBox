package channel

import (
  "fmt"
  "sync"
  "testing"
  "time"
)

func TestGorCounter(t *testing.T) {
  var mut sync.Mutex
  counter := 0
  for i := 0; i < 5000; i++ {
    go func() {
      //defer mut.Unlock()            // 1

      defer func() {                  // 2
        mut.Unlock()
      }()

      mut.Lock()
      counter++
    }()
  }
  time.Sleep(3 * time.Second)
  fmt.Println("counter: ", counter)
}

func TestGorCounterWait(t *testing.T) {
  var mut sync.Mutex
  counter := 0
  var wg sync.WaitGroup
  for i := 0; i < 5000; i++ {
    wg.Add(1)
    go func() {
      defer func() {
        mut.Unlock()
      }()

      mut.Lock()
      counter++
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println("counter: ", counter)
}










