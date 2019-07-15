package main

import (
  "fmt"
  "github.com/r00tjimmy/ColorfulRabbit"
  "sync"
  "time"
)

func moo(wg *sync.WaitGroup, i int){
  defer wg.Done()
  time.Sleep( 2 * time.Second)
  fmt.Println(i)
}

func main() {
  t1 := time.Now()
  var wg sync.WaitGroup

  for i := 0; i < 10; i++ {
    for j := 0; j < 20; j++ {
      wg.Add(1)
      go moo(&wg, j)
    }
  }

  wg.Wait()

  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}











