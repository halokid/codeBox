package main 

import (
  "fmt"
  "github.com/zeromq/goczmq"
)

func main() {
  push, err := goczmq.NewPush("tcp://127.0.0.1:31337")
  if err != nil {
    panic(err)
  }

  pull, err := goczmq.NewPull("tcp://127.0.0.1:31337")
  if err != nil {
    panic(err)
  }

  err = push.SendFrame([]byte("Hello World"), goczmq.FlagNone)
  if err != nil {
    panic(err)
  }

  frame, sz, err := pull.RecvFrame()
  if err != nil {
    panic(err)
  }

  fmt.Printf("We received a message of size %d\n", sz)
  fmt.Printf("The message was: '%s'\n", frame)

  pull.Destroy()
  push.Destroy()
}