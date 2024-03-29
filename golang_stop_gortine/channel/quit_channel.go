package main
/**
用 channel 来控制gor的退出
本质上是要用channel来调用 func 才可以
 */
import (
  "fmt"
  "time"
)

func doWork() int {
  fmt.Println("doWork ---------- ", 1)
  return 1
}

func main() {
  ch := make(chan int, 100)
  done := make(chan struct{})

  go func() {
    for {
      select {
      case ch <-doWork():
      case <-done:
        close(ch)     // 关闭上面才会停止监听，不然就会deadlock
        return
      }
      // 0.1秒执行一次  doWork, 一旦关闭 ch, 则停止执行
      time.Sleep(100 * time.Millisecond)
    }
  }()

  go func() {
    time.Sleep(2 * time.Second)
    done <-struct{}{}
  }()

  for i := range ch {
    fmt.Println("receive val: ", i)
  }

  fmt.Println("done...")

  time.Sleep(10 * time.Second)
}
