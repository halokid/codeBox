package main

import (

)
import (
  "os"
  "fmt"
  "net"
  "bytes"
  "io"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("xxxxx")
    os.Exit(1)
  }


  service := os.Args[1]

  conn, err := net.Dial("tcp", service)
  checkError(1, err)

  _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
  checkError(2, err)

  res, err := readFully(conn)
  checkError(3, err)

  fmt.Println(string(res))

  os.Exit(0)
}

func checkError(code int, err error) {
  if err != nil {
    fmt.Println(code)
    os.Exit(1)
  }
}



func readFully (conn net.Conn) ([]byte, error) {
  defer conn.Close()

  res := bytes.NewBuffer(nil)
  var buf [512]byte

  for {   // 循环读取服务器发来的信息
    n, err := conn.Read(buf[0:])    // 服务器每一次发来的信息是定量的，而不是全部的信息
    res.Write(buf[0:n])     // 把信息写入 res
    if err != nil {
      if err == io.EOF {
        break
      }
      return nil, err
    }
  }

  return res.Bytes(), nil
}




