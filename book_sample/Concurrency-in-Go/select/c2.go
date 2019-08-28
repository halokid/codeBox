package main

import (
  "fmt"
  "time"
)

func main() {
  start := time.Now()
  c := make(chan interface{})


  go func() {
    //close(c)
    time.Sleep(2 * time.Second)
    close(c)      // 注释这里就是死锁, 因为下面一直在读，但是却没有写入channel
  }()

  fmt.Println("block to read...")
  /**
  select {
  case <-c:   // 一直阻塞， 直到channel有东西可以读取， 或者close
    fmt.Println("unblocked %v later ", time.Since(start))
  }
  */


  for {     // 证明for的作用只是不断循环,  for 是不断执行select去监听的
    select {      // select 是不断监听 channel 有没有可以读取的数据， 注意不是监听变化，是监听有没有可以读取的数据
    case <-c:
      fmt.Println("加上for之后，可以对比一下 ", time.Since(start))
    }
  }

}



