package main

import "fmt"

/**
func main() {
  ch := make(chan int)
  ch <- 1
  fmt.Println("wont run here")
}
*/


var ch1 = make(chan int)
var ch2 = make(chan int)

func says(s string) {
  fmt.Println(s)
  ch1 <- <- ch2
}

func main() {
  go says("hello")
  <- ch1
}



