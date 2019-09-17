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
        //close(ch)
        return
      }
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
