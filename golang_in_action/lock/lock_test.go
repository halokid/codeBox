package lock

import (
  "fmt"
  "sync"
  "testing"
)

var cache map[string]string

const NUM_OF_READER = 40
const READ_TIMES = 100000

func init() {
  cache = make(map[string]string)

  cache["a"] = "aa"
  cache["b"] = "bb"
}

func lockFreeAccess() {
  var wg sync.WaitGroup
  wg.Add(NUM_OF_READER)
  for i := 0; i < NUM_OF_READER; i++ {
    go func() {
      for j := 0; j < READ_TIMES; j++ {
        _, err := cache["a"]
        if !err {
          fmt.Println("没有读错误")
        }
      }
      wg.Done()
    }()
  }
  wg.Wait()
}

func lockAccess()  {
  var wg sync.WaitGroup
  wg.Add(NUM_OF_READER)
  m := new(sync.RWMutex)
  for i := 0; i < NUM_OF_READER; i++ {
    go func() {
      for j := 0; j < READ_TIMES; j++ {
        m.RLock()
        _, err := cache["a"]
        if !err {
          fmt.Println("没有读错误")
        }
        m.RUnlock()
      }
      wg.Done()
    }()
  }
  wg.Wait()
}

func BenchmarkLockFree(b *testing.B) {
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    lockFreeAccess()
  }
}

func BenchmarkLock(b *testing.B) {
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    lockAccess()
  }
}













