package main

import "fmt"

func main() {
  /**
  stringStream := make(chan string)
  go func() {
    time.Sleep(2 * time.Second)
    stringStream <- "hello world"
  }()

  // 这里会一直阻塞， 只要其他的 channel 是逻辑正确的， 可以写入 channel， 这里都是一直在等待写入， 进行读取
  fmt.Println(<-stringStream)
  */


  /**
  // 死锁程序示范
  stringStream := make(chan string)
  go func() {
    if 0 != 1 {
      return
    }
    // 这里永远不会执行， 所以下面要取出就是死锁了
    stringStream <- "hello world"
  }()

  fmt.Println(<-stringStream)
  */


  stringStream := make(chan string)
  go func() {
    stringStream <- "hello world"
  }()

  // fixme: 如果这里执行， 上面的 go func channel 就是报错， 因为已经close了，上面还写入就报错
  close(stringStream)

  salutation, ok := <-stringStream
  fmt.Printf("(%v), %v", ok, salutation)
}

















