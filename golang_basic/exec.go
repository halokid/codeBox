package main 

import (
  "bytes"
  "fmt"
  "os/exec"
)

func main() {
  in := bytes.NewBuffer(nil)
  cmd := exec.Command("sh")
  cmd.Stdin = in

  go func() {
    in.WriteString("echo hello world > test.txt\n")
    //in.WriteString("ls\n")
    in.WriteString("exit\n")
  } ()

  if err := cmd.Run(); err != nil {
    fmt.Println(err)
    return
  }
}
