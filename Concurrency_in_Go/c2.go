package main

import (
  "fmt"
  "time"
)

func main() {
  bufferChan := make(chan string, 3)

  go func() {
    bufferChan <- "first"
    fmt.Println("Sent 1st")

    bufferChan <- "second"
    fmt.Println("Sent 2nd")

    bufferChan <- "third"
    fmt.Println("Sent 3rd")
  }()


  go func() {
    fmt.Println("Receiveing...")

    firstRead := <- bufferChan
    fmt.Println(firstRead)

    secondRead := <- bufferChan
    fmt.Println(secondRead)

    thirdRead := <- bufferChan
    fmt.Println(thirdRead)
  }()


  <-time.After(time.Second * 5)
}








