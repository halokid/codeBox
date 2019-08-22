package main
/**
golang装饰器模式
 */

import "fmt"

func userLogging(fun func()) func() { // 装饰器函数
  wrapper := func() {
    fmt.Println("this func is", fun)
    fun()
    fmt.Println("the end of foo")
  }
  return wrapper
}

func foo()  {
  println("i am foo")
}

func main() {
  foo := userLogging(foo)
  foo()
}

//this func is 0x490840
//the end of foo
//0x490840

