package main 

import (
  "fmt"
)

func main() {
  origin, wait := make(chan int), make(chan struct{})
  Proccessor(origin, wait)
  
  for num := 2, num < 10000; num++ {
    origin <- num
  }
  close(origin)
  <-wait
}


func Proccessor(seq chan int, wait chan struct{}) {
  go func() {
    prime, ok := <-seq
    if !ok {
      close(wait)
      return
    }
    fmt.Println(prime)
    
    out := make(chan int)
    Proccessor(out, wait)
    for num := range seq {
      if num % prime != 0 {
        out <- num
      }
    }
  }
}




// Wait 模式 -----------
wg := sync.WaitGroup{}
wg.Add(3)

go func() {
  defer wg.Done()
  // do ...
}



// Cancel 模式 ----------

import "context"

func Proc(ctx context.Context) {
  for {
    select {
      case <-ctx.Done():
        return
      default:
        // do ...
    }
  }
}



ctx := context.Backgroud()
ctx, cancel := context.WithCancel(ctx)
go Proc(ctx)
go Proc(ctx)
go Proc(ctx)

//cancle after 1s
time.Spleep(time.second)
cancle()




func Handler(r *Request) {
  timeout := r.Value("timeout")
  ctx, cancel := context.WithTimeout(context.Backgroud(), timeout)
  defer cancel()
  done := make(chan struct{}, 1)
  
  go func() {
    RPC(ctx, ...)
    done <- struct{}
  }()
  
  select{
    case <- done:
      // nice ...
    case <- ctx.Done():
      // timeout ...
  }
}



















































