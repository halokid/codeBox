package main

import (
  "fmt"
)

func recoverName() {
  if r := recover(); r!= nil {
    fmt.Println("recovered from ", r)
  }
}

func fullName(firstName *string, lastName *string) {
  defer recoverName()
  if firstName == nil {
    panic("runtime error: first name cannot be nil")
  }
  if lastName == nil {
    panic("runtime error: last name cannot be nil")
  }
  fmt.Printf("%s %s\n", *firstName, *lastName)
  fmt.Println("returned normally from fullName")
}

func main() {
  defer fmt.Println("deferred call in main")
  firstName := "Elon"
  fullName(&firstName, nil)
  fmt.Println("returned normally from main")
}




