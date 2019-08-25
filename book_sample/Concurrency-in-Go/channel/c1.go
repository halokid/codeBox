package main

import "fmt"

func main() {
  // 声明一个通用的channel
  var dataStream chan interface{}
  dataStream = make(chan interface{})
  fmt.Println(dataStream)

  // 声明一个只能被单向读取的channel
  var dataStreamx <-chan interface{}
  dataStreamx = make(<-chan interface{})
  fmt.Println(dataStreamx)

  // 声明一个只能被单向写入的channel
  var dataStreamy chan<- interface{}
  dataStreamy = make(chan<- interface{})
  fmt.Println(dataStreamy)

  var receiveChan <-chan interface{}
  var sendChan chan<- interface{}
  dataStreamk := make(chan interface{})

  // 这样做是有效的
  receiveChan = dataStreamk
  sendChan = dataStreamk
  fmt.Println(receiveChan)
  fmt.Println(sendChan)

}