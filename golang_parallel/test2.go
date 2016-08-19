package main

import (
    "fmt"
)

var a string
var c = make(chan int, 10)

func f() {
  fmt.Println("f 函数运行了")
  a = "hello world"
  c <- 1
}


func main() {
  fmt.Println("先让 f 函数先跑着")
  go f()
  <-c
  // a = "xx"
  fmt.Println("我这里取得了 a 的值", a)
}