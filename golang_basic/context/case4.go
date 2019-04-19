package main

import (
  "context"
  "fmt"
  "time"
)

var key string = "name"

func main() {
  valueCtx := context.WithValue(ctx, key, "【监控1】") 
  ctx, cancel := context.WithCancel(context.Background())

  go watch(valueCtx)
  time.Sleep(10 * time.Second)
  fmt.Println("可以了，通知监控停止")
  cancel()
  time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
  for {
    select {
    case <- ctx.Done():
      fmt.Println(ctx.Value(key), "监控退出，停止了...")
      return
    default:
      fmt.Println(ctx.Value(key), "goroutine监控中...")
      time.Sleep(2 * time.Second)
    }
  }
}




