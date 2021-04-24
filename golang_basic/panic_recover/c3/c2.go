package main

import "log"

func main() {
  defer log.Println("in main")

  defer func() {

    defer func() {
      panic("panic again and again")
    }()
    panic("panic again")
  }()

  panic("panic once")
}
