package main

import (
  "fmt"
  "sync"
)

func main() {
  var is []int

  for i := 0; i < 100; i++ {
    is = append(is, i)
  }
  fmt.Println(len(is))
  //os.Exit(0)


  var wg4 sync.WaitGroup
  for j := range is {
    wg4.Add(1)

    go func() {
      defer wg4.Done()
      fmt.Println(j)      // 当 fork 的 groutine 执行输出 j 的时候， main groutine 给 j 赋值是什么，就输出什么
    }()
  }

  wg4.Wait()
}
