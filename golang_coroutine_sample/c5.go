package main

import "fmt"

func main() {
  var ch = make(chan int, 4)
  for i := 0; i < cap(ch); i++ {
    ch <- i
  }

  for len(ch) > 0 {
    var value = <- ch
    fmt.Println(value)
  }
}



