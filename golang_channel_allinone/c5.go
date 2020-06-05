package main

import (
  "fmt"
)

func testQuit() {
  g := make(chan int)
  quit := make(chan bool)

  go func() {
    ///**
    for {       // for循环监听channel，channel写入多少，就读取多少，不存在写入一部分，读取结束的问题，这个就是for select的设计机制
      select {
      case v := <-g:
        fmt.Println(v)
      case <-quit:
        fmt.Println("退出")
        return    // 在for的select里面要用return退出
      }
    }
     //*/

     /**
    for {
      v, ok := <-g
      fmt.Println(v, ok)
      if !ok {        // 不close channel，ok就不会返回flase
        break
      }
    }
      */
  }()

  for i := 0; i < 100; i++ {
   g <-i
  }
  close(g)

  //time.Sleep(time.Second * 5)
  quit <-true
  fmt.Println("testQuit退出")
}

func main() {
  testQuit()
}
