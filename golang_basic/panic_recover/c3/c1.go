package main

import (
  "log"
)

func main() {
  defer log.Println("in main")
  if err := recover(); err != nil {
    log.Println(err)
  }

  panic("unknown err")
}




