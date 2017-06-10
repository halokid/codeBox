package main

import ()
import (
  "os"
  "fmt"
  "net"
  "bufio"
  "strings"
  "bytes"
)

const (
  uiDir = "dir"
  uiCd = "cd"
  uiPwd = "pwd"
  uiQuit = "quit"
)

const (
  DIR = "DIR"
  CD = "CD"
  PWD = "PWD"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("USAGE: ")
    os.Exit()
  }

  host := os.Args[1]

  conn, err := net.Dial("tcp", host + ":1202")
  checkError(1, err)

  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString("\n")

    line = strings.TrimRight(line, "\r\n")

    if err != nil {
      break
    }

    strs := strings.SplitN(line, " ", 2)

    switch strs[0] {
    case uiDir:
      dirRequest(conn)    //----------------
    case uiCd:
      if len(strs) != 2 {
        fmt.Println("cd <dir>")
        continue
      }
      fmt.Println("CD \"", strs[1], "\" ")
      cdRequest(conn, strs[1])      //-----------------
    case uiPwd:
      pwdRequest(conn)        //-----------------------
    case uiQuit:
      conn.Close()
      os.Exit(0)
    default:
      fmt.Println("unknow command")

    }

  }
}



func dirRequest(conn net.Conn) {
  //----------------
}

func cdRequest(conn net.Conn, dir string) {

}

func pwdRequest(conn net.Conn) {

}



































