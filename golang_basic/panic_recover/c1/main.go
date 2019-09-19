package main

import "fmt"

func r() int {
  if r := recover(); r != nil {
    return 1
  }
  return 0
}

func a() int {
  defer r()
  n := []int{1, 2, 3}
  fmt.Println(n[3])
  fmt.Println("a函数正常返回")
  return n[3]
}

func main() {
  a := a()
  fmt.Println("a-----", a)
  fmt.Println("main函数正常返回")
}
