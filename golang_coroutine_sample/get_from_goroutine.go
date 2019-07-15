package main

import (
  "fmt"
  "sync"
)

/**
从 goroutine 中返回数据
 */

func foo() string {
  wgx.Done()
  return "foo"
}

var wgx sync.WaitGroup

func main() {
  for i := 0; i < 10; i++ {
    wgx.Add(1)
    // 这个是错误的写法
    f := go foo()
    fmt.Println(f)
  }

  wgx.Wait()
}
