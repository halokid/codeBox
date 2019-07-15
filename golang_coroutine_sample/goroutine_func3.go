package main

import (
  "fmt"
  "github.com/r00tjimmy/ColorfulRabbit"
  "sync"
  "time"
)

func hoo(wg *sync.WaitGroup, i int){
  defer wg.Done()
  time.Sleep( 2 * time.Second)
  fmt.Println(i)
}

func gHoo() {
  // 这个函数的  wg 是放在循环外面
  var wg sync.WaitGroup
  for i := 0; i < 20; i++ {
    wg.Add(1)

    go hoo(&wg, i)
  }
  wg.Wait()
}

func main() {
  t1 := time.Now()

  for i := 0; i < 10; i++ {
    gHoo()
  }

  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}




















