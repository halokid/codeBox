package main

import (
  "fmt"
  "sync"
)

func main() {
  var wg2 sync.WaitGroup
  salutation := "hello3"

  wg2.Add(1)
  go func() {
    defer wg2.Done()
    salutation = "wellcome"
  }()
  wg2.Wait()
  fmt.Println(salutation)
}
