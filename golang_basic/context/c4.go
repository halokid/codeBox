package main

import (
  "fmt"
  "sync"
  "time"

  "golang.org/x/net/context"
)

var (
  wg sync.WaitGroup
)

func work(ctx context.Context) error {
  defer wg.Done()

  for i := 0; i < 1000; i++ {
    select {
    case <-time.After(2 * time.Second):
      fmt.Println("Doing some work ", i)

    // we received the signal of cancelation in this channel
    case <-ctx.Done():
      fmt.Println("Cancel the context ", i)
      return ctx.Err()
    }
  }
  return nil
}

func main() {
  // fixme: 定义一个上下文的属性， 然后传递这个上下文去影响其他的 goroutine， 这个是定义一个上下文， 去影响 work 的协程
  ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
  defer cancel()

  fmt.Println("Hey, I'm going to do some work")

  wg.Add(1)
  go work(ctx)
  wg.Wait()

  fmt.Println("Finished. I'm going home")
}




