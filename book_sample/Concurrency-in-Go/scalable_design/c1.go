package main

import (
  "fmt"
  "runtime/debug"
)

/**
错误传递
 */

type MyError struct {
  Inner           error
  Message         string
  StackTrace      string
  Misc            map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
  return MyError{
    Inner:              err,
    Message:            fmt.Sprintf(messagef, msgArgs...),
    StackTrace:         string(debug.Stack()),
    Misc:               make(map[string]interface{}),
  }
}


func main() {

}
