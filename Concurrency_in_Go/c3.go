package main

import (
  "fmt"
  "time"
)

func main() {
  theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
  oreChannel := make(chan string)
  minedOreChan := make(chan string)


  // finder
  go func(mine [5]string) {
    for _, item := range mine {
      if item == "ore" {
        oreChannel <- item      // send item on oreChannel
      }
    }
  }(theMine)


  // ore breaker
  go func() {
    for i := 0; i < 3; i++ {
      foundOre := <- oreChannel     // read fromm oreChannel
      fmt.Println("From Finder: ", foundOre)
      minedOreChan <- "minedOre"      // send to minedOreChan
    }
  }()

  // Smelter
  go func() {
    for i := 0; i < 3; i++ {
      minedOre := <- minedOreChan     // read from MinedOreChan
      fmt.Println("From Miner: ", minedOre)
      fmt.Println("From Smelter: Ore is smelted")
    }
  }()

  <-time.After(time.Second * 5)
}












