package main

import (
  "fmt"
  "math/rand"
)


/**
func rand_generator_1() int {
  return rand.Int()
}
**/


func rand_generator_2() chan int {
  out := make(chan int)
  
  go func() {
    for {
      out <- rand.Int()
      // fmt.Println("1111")
    }
  }()
  // fmt.Println()
  return out
}





func main() {
  rand_service_handler := rand_generator_2()
  fmt.Printf("%d\n", <-rand_service_handler)
}





































