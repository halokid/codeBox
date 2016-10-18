package main

import (
  "net"
  "time"
)

func main()  {
  service := ":1280"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    go handleClient(conn)
  }
}


/**
协程就是把 代码段封装起来， 然后并行地运行
 */

func handleClient(conn net.Conn) {
  defer conn.Close()
  daytime := time.Now().String()
  conn.Write([]byte(daytime))
  // finished with this client
}






