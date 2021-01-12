package main

/*
关于依赖注入的关键描述
1. 我们的函数不需要关心在哪里打印，以及如何打印，所以我们应该接收一个接口，而非一个具体的类型。

 */

import (
  "fmt"
  "io"
)

// 早期实现
//func Greet(writer *bytes.Buffer, name string) {
//  fmt.Fprintf(writer, "Hello, %s", name)
//}

// todo: 后期实现, 目的是为了也支持 os.stdout
// todo: fmt.Fprintf 允许传入一个 io.Writer 接口，我们知道 os.Stdout 和 bytes.Buffer 都实现了它
func Greet(writer io.Writer, name string) {
  fmt.Fprintf(writer, "Hello, %s", name)
}





