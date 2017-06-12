package main

import (
  "os"
  "fmt"
  "net/rpc"
  "log"
)

type Arith struct {
  A, B int
}


type Quotient struct {
  Quo, Rem int
}


func main() {
  if len(os.Args) != 2 {
    fmt.Println("usage: ", os.Args[0], "server")
    os.Exit(1)
  }

  serverAddress := os.Args[1]

  client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
  if err != nil {
    log.Fatal("diaing", err)
  }

  //call
  args := Arith{17, 8}  //其实这样的定义， args就是一个指针了
  // 如果是定义成一个结构体的话应该是  args2 := new(Args)
  var reply int
  err = client.Call("Arith.Multiply", args, &reply) // why args not point??
  if err != nil {
    log.Fatal("arith error: ", err)
  }
  fmt.Printf("arith: %d*%d = %d\n", args.A, args.B, reply)


  var quot Quotient
  err = client. Call( "Arith.Divide", args, &quot)
  if err != nil {
    log.Fatal("arith error", err)
  }
  fmt.Printf("arith: %d/%d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}












