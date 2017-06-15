package main

import (
  "fmt"
)

var ch1 chan int
var ch2 chan int

var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}


func main() {
  select {
    case getChan(0) <- getNumber(2):
      fm
  }
}