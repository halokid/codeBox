package main

import (
  "net"
  "os"
  "fmt"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Printf(os.Stderr, "Usage: %s host:port", os.Args[0])
    os.Exit(1)
  }

  service := os.Args[1]
  udpAddr, err := net.ResolveUDPAddr("udp4", service)
  checkError(err)
  conn, err := net.DialUDP("udp", nil, udpAddr)
  checkError(err)
  _, err = conn.Write([]byte("anything"))
  checkError(err)

  var buf [512]byte
  n, err := conn.Read(buf[0:])
  checkError(err)
  fmt.Println(string(buf[0:n]))
  os.Exit(0)
}


