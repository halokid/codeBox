package main

import (
  "fmt"
  "sync"
)

func main() {
  myPool := &sync.Pool{
    New: func() interface{} {
      fmt.Println("Creating new instance")
      return struct {}{}
    },
  }

  myPool.Get()    // 1
  instance := myPool.Get()      // 1
  // fixme: 使用完之后， 把实例放回 Pool 去
  myPool.Put(instance)          // 2
  myPool.Get()                  // 3
}
