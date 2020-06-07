package main

import (
  "fmt"
  "time"
)

/** 生产者 消费者 */

func testPCB() {
  fmt.Println("生产者和消费者范例")

  intChan := make(chan int)
  quitChan := make(chan bool)
  quitChan2 := make(chan bool)

  value := 0

  go func() {
    for i := 0; i < 3; i++ {
      value = value + 1
      intChan <-value
      fmt.Println("写入完成， value", value)
      time.Sleep(time.Second)
    }
    quitChan <-true
  }()

  go func() {
    for {
      select {
      case v := <-intChan:
        fmt.Println("读取完成， value", v)
      case <-quitChan:
        quitChan2 <-true
        return
      }
    }
  }()

  <-quitChan2
  fmt.Println("收到关闭信号，任务完成")
}

func main() {
  testPCB()
}

