package main 

import (
  "sync"
  "fmt"
)

func main() {
  var l *sync.RWMutex
  l = new(sync.RWMutex)
  l.RUnlock()
  fmt.Println("l")
  l.RLock()
}