package main
/**
协程挖矿程序的演示
 */

import (
  "fmt"
  "time"
)

func main() {
  theMine := [5]string{"ore1", "ore2", "ore3"}
  oreChan := make(chan string)

  // finder
  go func(mine [5]string) {
    for _, item := range mine {
      oreChan <- item     // send
    }
  }(theMine)

  // Ore Breaker
  go func() {
    for i := 0; i < 3; i++ {
      foundOre := <- oreChan    // receive
      fmt.Println("Miner: Received " + foundOre + " from finder")
    }
  }()

  <-time.After(time.Second * 5)   // again, ignore this for now
}



