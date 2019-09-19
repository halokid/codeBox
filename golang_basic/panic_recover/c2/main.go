package main

import (
  "fmt"
  "runtime/debug"
)

func takes(s []string, i int) string {
  defer func() string {
    if err := recover(); err != nil {
      debug.PrintStack()
      return "bug 1"
    }
    return "bug 2"
  }()
  return s[i]
}

func main() {
  sx := []string{"a", "b", "c"}
  c := takes(sx, 3)
  fmt.Println("c --------------", c)
}
