package main

import (
  "fmt"
  "runtime/debug"
)

func r() int {
  if r := recover(); r != nil {
    fmt.Println("111")
    debug.PrintStack()
    return 1
  } else {
    return 2
  }
  //return 0
}

func a() int {
  //defer r()
  n := []int{1, 2, 3}
  fmt.Println(n[3])
  fmt.Println("n[3] ------------", n[3])
  fmt.Println("a函数正常返回")
  return n[3]
}

func main() {
  a := a()
  fmt.Println("a-----", a)
  fmt.Println("main函数正常返回")
}
