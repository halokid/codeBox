package main

import "fmt"

func main() {
  var ch = make(chan int, 4)
  ch <- 1
  ch <- 2

  close(ch)

  value := <- ch
  fmt.Println(value)

  value = <- ch
  fmt.Println(value)

  value = <- ch
  fmt.Println(value)
}
