package main

import (
  "errors"
  "net/rpc"
  "net"
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
    return errors.New("divide by zero")
  }

  quo.Quo = args.A / args.B
  quo.Rem = args.A % args.B
  return nil
}

func main()  {
  arith := new(Arith)
  rpc.Register(arith)

  tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
  checkError(1, err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(2, err)


  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    rpc.ServeConn(conn)
  }

}
























