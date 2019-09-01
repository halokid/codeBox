package main

import (
  "fmt"
  "sync"
)

/**
对比channel通信 和 指定内存 通信的性能
此为 指定内存 的方式
此程序为同步陷阱的范例
 */

func main() {
  setMem := make(map[int]int)
  wg := sync.WaitGroup{}
  lk := sync.RWMutex{}

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(wg *sync.WaitGroup) {
      fmt.Println("i ---- ", i)
      defer wg.Done()
      lk.Lock()
      setMem[i] = i * 2
      lk.Unlock()
    }(&wg)
  }

  wg.Wait()
  fmt.Println("setMen len ----------", len(setMem))
  fmt.Println("setMem done ... ")
}
