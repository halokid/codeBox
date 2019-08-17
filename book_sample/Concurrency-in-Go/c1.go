package main

import (
  "fmt"
  "time"
)

func main() {
  go sayHello()
  // contine doing other things

  go func() {
    fmt.Println("hello 1")
  }()

  sayHello := func() {
    fmt.Println("hello 2")
  }
  go sayHello()

  time.Sleep(2 * time.Second)
}

func sayHello()  {
  fmt.Println("hello")
}
