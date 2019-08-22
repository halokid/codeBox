package main

import (
  "fmt"
  "reflect"
)

func Decorator(decoPtr, fn interface{}) (err error) {
  var decoratedFunc, targetFunc reflect.Value

  decoratedFunc = reflect.ValueOf(decoPtr).Elem()
  targetFunc = reflect.ValueOf(fn)

  v := reflect.MakeFunc(targetFunc.Type(),
    func(in []reflect.Value) (out []reflect.Value) {
      fmt.Println("before")
      args := []reflect.Value{}
      if len(in) == 1 && in[0].Kind() == reflect.Slice {
        for i := 0; i < in[0].Len(); i++ {
          args = append(args, in[0].Index(i))
        }
        in = args
      }
      out = targetFunc.Call(in)
      fmt.Println("after")
      return
    })

  decoratedFunc.Set(v)
  return
}

func foox(a, b, c int) int {
  fmt.Println(a)
  return b
}


func main() {
  myfoo := foox
  Decorator(&myfoo, foox)
  myfoo(1, 2, 3)
}
//before
//1
//after

