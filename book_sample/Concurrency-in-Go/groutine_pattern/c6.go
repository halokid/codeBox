package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  newRandStream := func(done <-chan interface{}) <-chan int {
    randStream := make(chan int)
    go func() {
      // 创建 运行时的这个 gor
      defer fmt.Println("newRandStream close exited")
      defer close(randStream)

      for {
        select {
        case randStream <- rand.Int():
        case <-done:      // 接收到done， 退出， 也在这个内存空间里（gor里）销毁这个 gor
          return
        }
      }
    }()

    return randStream
  }

  done := make(chan interface{})        // 在 main 这个 gor 创建 done 这个 channel
  randStream := newRandStream(done)     // 发送done chanel
  fmt.Println("3 random ints: ")
  for i := 1; i <= 3; i++ {
    fmt.Println(i, ":", <-randStream)
  }

  close(done)                           // 也在 mian 这个gor 销毁 done 这个 channel
  time.Sleep(2 * time.Second)
}




