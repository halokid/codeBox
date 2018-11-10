package main 

/**

defer 是后进先出。
panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。

**/


import (
  "fmt"
) 

func main() {
  deferCall()
}


func deferCall() {
  defer func() {
    fmt.Println("before print")
  }()

  defer func() {
    fmt.Println("in print") 
  }()

  defer func() {
    fmt.Println("after print")
  }()

  panic("trigger panic")
}