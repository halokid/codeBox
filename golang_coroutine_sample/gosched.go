package main

import (
  "fmt"
  "runtime"
)

func say(s string)  {
  //runtime.GOMAXPROCS(2)
  for i := 0; i < 1000000000000; i++ {
    runtime.Gosched()     // 显式的让出CPU时间，让其他的 goroutine 运行
    fmt.Printf("%d----------%s\n", i, s)
  }
}

func main() {
  // 真正达到多核的能力，必须要加上这个声明，如果没有这个声明，则 goroutine 则都是跑在单个线程里面（就是多个单独的线程，对应处理更多个goroutine）
  // 多个  gorontine 互相抢占线程
  //runtime.GOMAXPROCS(4)
  go say("world")
  say("hello")
}



