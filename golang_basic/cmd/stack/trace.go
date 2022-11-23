package main

import (
  "log"
  "runtime"
)

/*
// TODO: version 1
func Trace(name string) func() {
  log.Println("enter: ", name)
  return func() {
    log.Println("exit: ", name)
  }
}

func foo() {
  defer Trace("foo")()
  bar()
}

func bar() {
  defer Trace("bar")()
}

func main() {
  defer Trace("main")()
  foo()
}
 */

 // -----------------------------------------------------------------

// TODO: version 2, 通过 Go 中的 runtime.Caller 获取到函数调用栈上的函数调用者信息（Caller(0) 获取函数信息，Caller(1) 获取函数调用者信息）
func Trace() func() {
  pc, _, _, ok := runtime.Caller(1)
  if !ok {
    panic(any("not found caller"))
  }

  fn := runtime.FuncForPC(pc)
  name := fn.Name()

  log.Println("enter: ", name)
  return func() {
    log.Println("exit: ", name)
  }
}

func foo() {
  defer Trace()()
  bar()
}

func bar() {
  defer Trace()()
}

func main() {
  defer Trace()()
  foo()
}

// -----------------------------------------------------------------










