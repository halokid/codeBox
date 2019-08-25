package main

import "fmt"

/**
用channel 比 sync.WaitGroup 性能好的原因之一就是
channel 不需要 wait 协程（从底层来说不需要监听各个协程的信号，状态等操作）
 */

func main() {
  intStream := make(chan int)

  go func() {
    defer close(intStream)
    for i := 1; i <= 5; i++ {
      intStream <- i
    }
  }()

  for it := range intStream {
    fmt.Println(it)
  }
}
