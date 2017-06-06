package main

import (
  "fmt"
  "net"
  "os"
  "time"
)


func main() {
  service := ":1200"
  udpAddr, err := net.ResolveUDPAddr("udp4", service)     //change to UDP addr
  checkError(err)

  conn, err := net.ListenUDP("udp", udpAddr)        // listen addr & get conn from client
  checkError(err)


  for {
    handleClient(conn)        // handle client conn
  }
}



func handleClient(conn *net.UDPConn) {
  var buf [512]byte

  _, addr, err := conn.ReadFromUDP(buf[0:])       // read from client
  if err != nil {
    return
  }

  daytime := time.Now().String()

  conn.WriteToUDP([]byte(daytime), addr)        // write to udp client
}

func checkError(err error) {
  if err != nil {
    fmt.Println("xxxxx")
    os.Exit(1)
  }
}
















