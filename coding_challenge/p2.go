package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
  // write your code in Go 1.4
  aLen := len(A)
  for _, item := range A {
    if item > aLen - 1 {
      return aLen - 1
    }
    return aLen + 1
  }
}
