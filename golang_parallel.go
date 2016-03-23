package main

import (
  "fmt"
  // "time"
  "runtime"
)

/**
var quit chan int = make(chan int)

func loop() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d ", i)
  }
  quit <- 0
}

func main() {
  go loop()
  go loop()
  
  for i := 0; i < 2; i++ {
    <- quit
  }
}
**/

/*8
var quit chan int

func foo(id int) {
  fmt.Println(id)
  time.Sleep(time.Second)
  quit <- 0
}

func main() {
  count := 1000
  quit = make(chan int, count)
  
  for i := 0; i < count; i++ {
    go foo(i)
  }
  
  for i := 0; i < count; i++ {
    <- quit
  }
}
**/

var quit chan int = make(chan int)

func loop() {
  for i := 0; i < 100; i++ {
    runtime.Gosched()
    fmt.Printf("%d ", i)
  }
  quit <- 0
}

func main() {
  runtime.GOMAXPROCS(2)
  
  go loop()
  go loop()
  
  for i := 0; i < 2; i++ {
    <- quit
  }
}























