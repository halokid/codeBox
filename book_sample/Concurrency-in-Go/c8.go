package main

import (
  "fmt"
  "sync"
)

func main() {
  type Button struct {
    Clicked   *sync.Cond
  }

  button := Button{
    Clicked:      sync.NewCond(&sync.Mutex{}),
  }

  subscribe := func(c *sync.Cond, fn func()) {
    var tempwd sync.WaitGroup
    tempwd.Add(1)
    go func() {
      tempwd.Done()   // fixme: 表示要走到这里才 done， 但是下面的执行不需要等待，因为下面的逻辑本身也是lock住了
      c.L.Lock()
      defer c.L.Unlock()
      c.Wait()
      fn()
    }()
    tempwd.Wait()
  }

  var wg sync.WaitGroup //3
  wg.Add(3)
  subscribe(button.Clicked, func() { //4
    fmt.Println("Maximizing window.")
    wg.Done()
  })
  subscribe(button.Clicked, func() { //5
    fmt.Println("Displaying annoying dialog box!")
    wg.Done()
  })
  subscribe(button.Clicked, func() { //6
    fmt.Println("Mouse clicked.")
    wg.Done()
  })

  button.Clicked.Broadcast() //7

  wg.Wait()
}




