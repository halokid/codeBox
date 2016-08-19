package main

import (
    "fmt"
    "runtime"
    "time"
)


var a int  = 1

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  
  go sheep(1)
  
  time.Sleep(time.Millisecond)
  fmt.Println("end", a)
}


func sheep(i int) {
  for ; ; i += 1 {
    fmt.Println(i, "个屌丝")
    a += 1
  }
}

