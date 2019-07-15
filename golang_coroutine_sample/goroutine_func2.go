package main

import (
  "fmt"
  "github.com/r00tjimmy/ColorfulRabbit"
  "sync"
  "time"
)

func poo(wg *sync.WaitGroup, i int){
  defer wg.Done()
  time.Sleep( 2 * time.Second)
  fmt.Println(i)
}

func gPoo() {
  // 这个函数的  wg 是放在循环里面
  for i := 0; i < 20; i++ {
    var wg sync.WaitGroup
    wg.Add(1)

    go poo(&wg, i)

    wg.Wait()
  }
}

func main() {
  t1 := time.Now()
  gPoo()
  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}


