package main 

import (
  "sync"
  "fmt"
)

/**
func main() {
  var l *sync.Mutex
  l = new(sync.Mutex)
  l.Unlock()
  fmt.Println("l")
  l.Lock()
}
**/


func main() {
  var l *sync.Mutex
  l = new(sync.Mutex)
  l.Lock()
  fmt.Println("l")
  l.Lock()
}