package main 

import (
  "fmt"
  "runtime"
  "sync"
) 


func main() {
  runtime.GOMAXPROCS(1)
  wg := sync.WaitGroup{} 
  wg.Add(20)

  for i := 0; i < 5000; i++ {
    go func() {
      fmt.Println("A: ", i)
      wg.Done()
    }()
  }


  wg.Wait() 
}








