package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  ctx, cancel := context.WithCancel(context.Background())

  valueCtx := context.WithValue(ctx, key, "add value")

  go watch(valueCtx)

  time.Sleep(10 * time.Second)
  cancel()

  time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
  for {
    select {
    case <- ctx.Done():
      fmt.Println(ctx.Value(key), "is cancel")
      return

    default:
      fmt.Println(ctx.Value(key), " int goroutine")
      time.Sleep(2 * time.Second)
    }
  }
}





