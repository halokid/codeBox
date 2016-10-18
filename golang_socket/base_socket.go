package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  service := "7777"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    // process code

    conn.Close()
  }
}

func checkError(err error) {
  if err != nil {
    fmt.Printf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
