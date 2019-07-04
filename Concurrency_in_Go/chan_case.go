package main

import "fmt"

var xx int = 0
var cx = make(chan int, 1)


func thread1(){
  <-cx // Grab the ticket
  xx++
  cx <- 1 // Give it back
}

func thread2(){
  <-cx
  xx++
  cx <- 1
}

func main() {
  cx <- 1 // Put the initial value into the channel
  go thread1()
  go thread2()
  fmt.Println(xx)
}

