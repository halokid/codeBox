package main

import (
  "fmt"
  "time"
)

func main() {
  tick := time.Tick(10)
  tick2 := time.Tick(20)
  tick3 := time.Tick(40)

  for  {
    select {
    case <- tick:
      fmt.Println("tick .")
    case <- tick2:
      fmt.Println("tick 2 .")
    case <- tick3:
      fmt.Println("tick 3 .")
    default:
      fmt.Println("      . ")
      time.Sleep(30)
    }
  }
}
