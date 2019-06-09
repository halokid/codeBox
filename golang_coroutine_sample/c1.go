package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("run in main goroutine")
  go func() {
    fmt.Println("run in child goroutine")
    go func() {
      fmt.Println("run in grand child goroutine")
      go func() {
        fmt.Println("run in grand grand child goroutine")
      }()
    }()
  }()

  time.Sleep(time.Second * 2)
  fmt.Println("main goroutine will quit")
}
