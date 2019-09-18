package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  forever := make(chan struct{})
  ctx, cancel := context.WithCancel(context.Background())

  go func(ctx context.Context) {
    for {
      select {
      case <-ctx.Done():
        forever <-struct{}{}
        return
      default:
        fmt.Println("for loop...")
      }
      time.Sleep(500 * time.Millisecond)
    }
  }(ctx)

  go func() {
    fmt.Println(" 3秒之后 cancel 掉 context，才写入  ctx.Done()")
    time.Sleep(3 * time.Second)
    cancel()
  }()

  // 阻塞的作用
  <-forever
  fmt.Println("done...")
}
