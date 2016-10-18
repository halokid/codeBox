package main

import (
  "net"
  "time"
  "strconv"
  "strings"
  "fmt"
)

func main()  {
  service := ":1200"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    go handleClientx(conn)
  }
}


func handleClientx(conn net.Conn) {
  conn.SetReadDeadline(time.Now().Add(2 * time.Minute))  //set 2 minutes timeout
  request := make([]byte, 128)  //set maxium request length to 128B to prevent flood attack
  defer conn.Close()      //close connection before exit

  for {
    read_len, err := conn.Read(request)   //重新读 request 变量

    if err != nil {
      fmt.Println(err)
      break
    }

    if read_len == 0 {          // 如果没收到心跳包，就断连接
      break     //connection already closed by client
    } else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {      // 时间戳的形式
      daytime := strconv.FormatInt(time.Now().Unix(), 10)
      conn.Write([]byte(daytime))
    } else {                                                                      // 时间字符串形式
      daytime := time.Now().String()
      conn.Write([]byte(daytime))
    }
    request = make([]byte, 128)       // clear last read content, 清空 request 变量

  } //END FOR
}




















