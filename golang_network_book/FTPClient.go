package main

import (

)
import (
  "os"
  "fmt"
  "net"
  "bufio"
  "strings"
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

  conn, err := net.Dial("tcp", host+":1202")
  checkError(1, err)


  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString("\n")

    line = strings.TrimRight(line, "\r\n")

    if err != nil {
      break
    }

    strs := strings.SplitN()

  }
}
