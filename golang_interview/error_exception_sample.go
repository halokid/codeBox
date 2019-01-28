package main 

/**
golang异常，错误处理代码
**/


/**
整个 异常错误 处理代码的逻辑如下:
1.  执行 test()， 先跑到 funcA()
2.  funcA() 先执行  funcB()
3.  funcB() 异常 panic
4.  触发  funcA  的  recover(),  p 捕获到的 panic 信息为 foo, 程序正常恢复
5.  最后 funcA() 没有返回 error， 所以程序正常执行

**/


import (
  "fmt"
  "errors"
  // "debug"
)

func funcA() error {
  defer func() { 
    if p := recover(); p != nil {
      fmt.Printf("panic recoverd, p: %v \n", p)
      // debug.PrintStack()
    }
  }()
  return funcB()
}


func funcB() error {
  // simulation 
  panic("foo")
  return errors.New("success")
}

func test() {
  err := funcA()
  if err == nil {
    fmt.Printf("err is nil \n")
  } else {
    fmt.Printf("err is %v \n", err)
  }
}


func main() {
  test() 
}








