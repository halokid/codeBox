package golang_rpc

import "net"

type Session struct {
  conn net.Conn
}

func NewSession(conn net.Conn) *Session {
  return &Session{conn:  conn}
}


