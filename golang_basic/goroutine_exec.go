package main

import (
  "fmt"
  "os/exec"
  "runtime"
  "time"
)

var quit chan int = make(chan int)

func runComm() {
  cmd := exec.Command("date", "+%s.%N")
  out, err := cmd.CombinedOutput()

  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(out))
  //return string(out)
  time.Sleep(3*time.Second)
  quit <- 0
}

func main() {
  runtime.GOMAXPROCS(2)
  
  go runComm()
  go runComm()

  for i := 0; i < 2; i++ {
    <- quit 
  }
}
