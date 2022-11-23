package main

import (
  "fmt"
  "github.com/go-stack/stack"
  "log"
)

func DoTheThing() {
  c := stack.Caller(0)
  log.Print(c)          // "source.go:10"
  log.Printf("%+v", c)  // "pkg/path/source.go:10"
  log.Printf("%n", c)   // "DoTheThing"

  s := stack.Trace().TrimRuntime()
  log.Print(s)          // "[source.go:15 caller.go:42 main.go:14]"
}

func Example_callFormat() {
  //logCaller("%+s")
  logCaller("%v   %[1]n()")
  // Output:
  // github.com/go-stack/stack/format_test.go
  // format_test.go:13   Example_callFormat()
  subTest()
}

func logCaller(format string) {
  fmt.Printf( format+"\n", stack.Caller(1))
}

func subTest() {
  //log.Println("this is subTest")
  logCaller("%v   %[1]n()")
}

func main() {
  DoTheThing()
  Example_callFormat()
}



