package main

import (
  "fmt"
  "github.com/r00tjimmy/ColorfulRabbit"
  "sync"
  "time"
)

func koo(wg *sync.WaitGroup, i int){
  defer wg.Done()
  time.Sleep( 2 * time.Second)
  fmt.Println(i)
}


func gKoo() {
  // 这个函数的  wg 是放在循环外面
  var wg sync.WaitGroup
  for i := 0; i < 20; i++ {
    wg.Add(1)

    go koo(&wg, i)
  }
  wg.Wait()
}

func main() {
  t1 := time.Now()
  gKoo()
  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}




