package main

import (
  "fmt"
  "sync"
)

func main() {
  var wg sync.WaitGroup

  sayHello2 := func() {
    defer wg.Done()
    fmt.Println("hello2")
  }

  wg.Add(1)
  go sayHello2()
  wg.Wait()   // 在这里加入连接点
}
