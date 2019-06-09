package main

import "fmt"

func sendx(ch chan int) {
  i := 0
  for {
    i++
    ch <- i
  }
}

func recvx(ch chan int) {
  value := <- ch
  fmt.Println(value)

  value = <- ch
  fmt.Println(value)

  close(ch)
}

func main() {
  var ch = make(chan int, 4)
  go recvx(ch)
  sendx(ch)
}






















