package main

import (
  "fmt"
  "sync"
)

var global int = 0
var m sync.Mutex
var wgxx sync.WaitGroup

func Thread1() {
  defer wgxx.Done()
  m.Lock()
  global = 1
  m.Unlock()
}

func Thread2() {
  defer wgxx.Done()
  m.Lock()
  global = 2
  m.Unlock()
}

func main() {
  wgxx.Add(1)
  go Thread1()
  fmt.Println(global)

  wgxx.Add(1)
  go Thread2()
  fmt.Println(global)
  wgxx.Wait()
  fmt.Println(global)
}

