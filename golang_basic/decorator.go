package main

import (
  "fmt"
  "time"
)

/**
一个装时期模式的例子
 */

func timeSpent(inner func(op int) int) func(op int) int {
  // 输入参数是一个函数， 输出也是一个函数
  return func(n int) int {
    start := time.Now()
    ret := inner(n)     // 计算这个函数的耗时
    fmt.Println("time spent: ", time.Since(start).Seconds())
    return ret
  }
}

func slowFunc(op int) int {
  time.Sleep(1 * time.Second)
  return op
}

func main() {
  tf := timeSpent(slowFunc)
  i := tf(10)
  fmt.Println(i)
}








