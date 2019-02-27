package golang_rpc

import (
  "fmt"
  "net"
  "reflect"
)

type Server struct {
  addr      string
  funcs     map[string]reflect.Value
}


func NewServer(addr string) *Server {
  return &Server{addr:  addr, funcs:  make(map[string]reflect.Value)}
}


func (s *Server) Run() {
  l, err := net.Listen("tcp", s.addr)
  if err != nil {
    fmt.Printf("listen on %s err: %v", s.addr, err)
  }

  for {
    conn, err := l.Accept()
    if err != nil {
      fmt.Printf("accept err: %v", err)
    }

    // todo
    srvSession := NewSession(conn)

    // read RPC call
    b, err := srvSession.conn.Read()
    if err != nil {
      fmt.Printf("read err: %v", err)
      return
    }

    // decode RPC call
    rpcData, err := decode(b)
    if err != nil {
      fmt.Printf("decode err: %v", err)
      return
    }

    // find RPC call func
    f, ok := s.funcs[rpcData.Name]
    if !ok {
      fmt.Printf("func %s not exists", rpcData.Name)
      return
    }

    // make RPC call func args
    inArgs := make([]reflect.Value, 0, len(rpcData.Args))
    for _, arg := range rpcData.Args {
      inArgs = append(inArgs, reflect.ValueOf(arg))
    }

    // call & run func
    out := f.Call(inArgs)
  }
}







