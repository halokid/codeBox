package main

import (
  "fmt"
  "time"
)

func main() {
  go say("world")
  say("hello")
  
  fmt.Println("---------------1")
  
  a := []int{7, 2, 8, -9, 4, 0}
  
  c := make(chan int)
  go sum(a[:len(a)/2], c)
  go sum(a[len(a)/2:], c)
  
  x, y := <-c, <-c
  
  fmt.Println(x, y, x+y)
  
  
  
  fmt.Println("----------------2")
  c2 := make(chan int)
  c2 <- 1
  c2 <- 2
  fmt.Println(<-c2)
  fmt.Println(<-c2)
  
  
  
  fmt.Println("-----------------3")
  c3 := make(chan int, 10)
  go fibonacci(cap(c3), c3)
  

  
}



