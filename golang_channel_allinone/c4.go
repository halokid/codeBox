package main

import (
  "fmt"
  "time"
)

func testMergeInput() {
  /** 将多个输入的channel进行合并成一个channel */
  input1 := make(chan int)
  input2 := make(chan int)
  output := make(chan int)

  go func(in1, in2 <-chan int, out chan<- int) {
    /** in1, in2接收的是chan，它们本身也是chan int类型。  out是一个chan，接收的是int */
    for {
      select {
      case v := <-in1:
        out <-v
      case v := <-in2:
        out <-v
      }
    }
  }(input1, input2, output)

  go func() {
    for i := 0; i < 10; i++ {
      input1 <-i
      time.Sleep(time.Millisecond * 100)
    }
  }()

  go func() {
    for i := 20; i < 30; i++ {
      input2 <-i
      time.Sleep(time.Millisecond * 100)
    }
  }()

  go func() {
    for {
      select {
      case value := <-output:
        fmt.Println("输出:", value)
      }
    }
  }()

}

func main() {
  testMergeInput()
  time.Sleep(time.Second * 5)
  fmt.Println("主协程退出")
}






