package main

import (
  "net"
  "os"
  "fmt"
)


func main() {
  if len(os.Args) != 2 {
    fmt.Println("need usage")
    os.Exit(1)
  }


  service := os.Args[1]
  udpAddr, err := net.ResolveUDPAddr("udp4", service)  //change to UDP addr
  checkError(err)

  conn, err := net.DialUDP("udp", nil, udpAddr)     // conn to server
  checkError(err)

  _, err = conn.Write([]byte("anything"))       // write things to server
  checkError(err)


  var buf [512]byte
  n, err := conn.Read(buf[0:])        // read response from server
  checkError(err)

  fmt.Println(string(buf[0:n]))

  os.Exit(0)

}


func checkError(err error) {
  if err != nil {
    fmt.Println("xxxx")
    os.Exit(1)
  }
}
















