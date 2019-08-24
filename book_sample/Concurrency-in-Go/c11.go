package main

import (
  "fmt"
  "sync"
)

func main() {
  var numCalcsCreated int
  calcPool := &sync.Pool{
    New: func() interface{} {
      numCalcsCreated += 1
      mem := make([]byte, 1024)
      return &mem
    },
  }

  // 将池扩充到 4KB
  calcPool.Put(calcPool.New())
  calcPool.Put(calcPool.New())
  calcPool.Put(calcPool.New())
  calcPool.Put(calcPool.New())

  const numWorkers = 1024 * 1024
  var wg sync.WaitGroup
  wg.Add(numWorkers)

  for i := numWorkers; i > 0; i-- {
    go func() {
      defer wg.Done()

      //mem := calcPool.Get().([]byte)
      // fxime: 为了效率， 实例返回的是指针而不是一块内存
      mem := calcPool.Get().(*[]byte)
      defer calcPool.Put(mem)
    }()
  }

  wg.Wait()
  fmt.Println(numCalcsCreated, " calculators where created ")
  fmt.Println("一共需要创建", numCalcsCreated, "个计算实例，池的最大需求容量为", numCalcsCreated)
}











