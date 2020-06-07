package main
/*
指定channel是输入还是输出型的，防止编写时写错误输入输出，指定了的话，可以在编译时期作错误的检查
*/

import (
  "fmt"
  "time"
)

func TestInAndOutChan() {
  ch := make(chan int)
  quit := make(chan bool)

  //输入型的chan是这种格式的：inChan chan<- int，如果换成输出型的，则编译时会报错
  go func(inChan chan <-int) {      // inChan为输入型
    for i := 0; i < 10; i++ {
      inChan <-i      // 接收输入
      time.Sleep(time.Millisecond * 500)
    }
    quit <-true
    quit <-true
  }(ch)

  go func(outChan <-chan int) {     // outChan为输出型
    for {
      select {
      case v := <-outChan:        // 输出到v
        fmt.Println("输出value", v)
      case <-quit:
        fmt.Println("收到退出通知， 退出")
        return
      }
    }
  }(ch)

  <-quit
  fmt.Println("收到退出通知，主协程退出")
}

func main() {
  TestInAndOutChan()
}

