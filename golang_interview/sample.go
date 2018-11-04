package main 

import (
  "fmt"
) 


func main() {
  if (true) {
    defer fmt.Println("1")
  } else {
    defer fmt.Println("2")
  }

  fmt.Println("3")


  s := make([]int, 5, 10)
  s[0] = 100
  fmt.Println(s)
}