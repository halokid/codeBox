package main

import (
  "fmt"
  "time"
)

func main() {
  timeout := make(chan bool, 1)
  go func() {
    fmt.Println("------------ 子进程1--------------")
    t1 := time.Now().UnixNano()
    fmt.Println(t1)
    fmt.Println("这个一定会执行")
    
    time.Sleep(3 * time.Second)
    // timeout <- true
  }()
  
  fmt.Println("首先逻辑还是响应 main 函数")
  
  go func() {
    fmt.Println("------------ 子进程2--------------")
    t2 := time.Now().UnixNano()
    fmt.Println(t2)
    fmt.Println("相当于fork一个子进程在进行")
  }()
  
  ch := make(chan int)
  select {
    case <-ch:
    case <-timeout:
      fmt.Println("------------ 回到main函数 --------------")
      fmt.Println("task is timeout!")
  }
  
  fmt.Println("main 函数本身的输出")
}



































