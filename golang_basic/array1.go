package main 

import (
  "fmt"
)

func main() {
  var n [10]int
  var i, j int
  
  for i = 0; i < 10; i++ {
    n[i] = i + 100
  }
  
  for j = 0; j < 10; j++ {
    fmt.Printf("Element[%d] = %d\n", j, n[j])
  }
}