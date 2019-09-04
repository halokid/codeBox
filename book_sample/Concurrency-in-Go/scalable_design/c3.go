package main

import (
  "fmt"
  "time"
)

func doWork()  {
  fmt.Println("doWork...")
}

func main() {
  t := time.AfterFunc(3 * time.Second, doWork)
  fmt.Println("t ------ ", t)
  //time.Sleep(5 * time.Second)
}
