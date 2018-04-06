package main

import (
  "sync"
  "context"
  "time"
)


// --------------------- WAIT 模式 -------------------------
func main() {
  wg := sync.WaitGroup{}
  wg.Add(3)

  go func() {
    defer wg.Done()

    //do...
  }()

  go func() {
    defer wg.Done()

    //do...
  }()

  go func() {
   defer wg.Done()

    //do...
  }()

//--------------------------- CANCLE 模式 -------------------------

  ctx := context.Background()
  ctx, cancle := context.WithCancel(ctx)
  go Proc(ctx)
  go Proc(ctx)
  go Proc(ctx)

  //cancle after 1s
  time.Sleep(time.Second)
  cancle()
}


func Proc(ctx context.Context) {
  for {
    select {
    case <-ctx.Done():
      return
    default:
      //do... 取消处理逻辑
    }
  }
}




























