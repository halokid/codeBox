package main

import (
  "os"
  "fmt"
  "net"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("usage");
    os.Exit(1)
  }

  addr, err := net.ResolveIPAddr("ip", os.Args[1])    //change network addr
  checkError(1, err)

  conn, err := net.DialIP("ip4:icmp", addr, addr)   //conn ip addr
  checkError(2, err)

  //define communication standard
  var msg [512]byte
  msg[0] = 8    //echo
  msg[1] = 0    //code 0
  msg[2] = 0    //checksum, fix later
  msg[3] = 0    //checksum, fix later
  msg[4] = 0    // identifier[0]
  msg[5] = 13    // identifier[1]
  msg[6] = 0    // sequence[0]
  msg[7] = 37    // sequence[1]
  len := 8

  check := checkSum(msg[0:len])
  msg[2] = byte(check >> 8)
  msg[3] = byte(check & 255)

}


func checkError(code int, err error) {
  if err != nil {
    fmt.Println(code)
    os.Exit(1)
  }
}
