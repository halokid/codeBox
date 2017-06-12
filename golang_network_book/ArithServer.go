package main

import (
  "errors"
  "net/rpc"
  "net/http"
  "fmt"
)

type Args struct {
  A, B int
}

type Quotient struct {
  Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
  *reply = args.A * args.B
  return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
  if args.B == 0 {
    return errors.New("drivide by zero")
  }

  quo.Quo = args.A / args.B
  quo.Rem = args.A % args.B

  return nil
}



func main() {
  arith := new(Arith)

  //golang已经集成了RPC的方法了
  rpc.Register(arith)
  rpc.HandleHTTP()

  err := http.ListenAndServe(":1234", nil)
  if err != nil {
    fmt.Println(err.Error())
  }

}






















