package main

import (
  "fmt"
  "time"
)

func chDoWork(hasDoneWork chan int) {
  time.Sleep(5 * time.Millisecond)
  fmt.Println("chDoWork...")
  hasDoneWork <- 1
}

func readDoneWork(hasDoneWork chan int, allDone chan int) {
  //defer close(hasDoneWork)
  for {
    select {
    case <-hasDoneWork:
      fmt.Println("read done work...")
    case <-allDone:
      fmt.Println("all done...")
      return
    }
  }
}

func checkDone(allDone chan int) {
  select {
  case <-allDone:
    fmt.Println("checkDone all done...")
    return
  }
}

func setDone(allDone chan int) {
  allDone <-1
}

func main() {
  hasDoneWork := make(chan int)
  workCount := 0

  go chDoWork(hasDoneWork)
  //workCount ++

  go chDoWork(hasDoneWork)
  //workCount++

  go chDoWork(hasDoneWork)
  //workCount++

  allDone := make(chan int)

  //readDoneWork(hasDoneWork)
  //allDone <-1     // 如果写在这里， readDoneWork这个gor是不会读取到 allDone 这个channel的， 因为这个跟下面的调用 readDoneWork 是同一个gor

  //allDone <-1
  //go checkDone(allDone)

  go setDone(allDone)

  // 这个for循环在这里是没用的， 因为开了一个新的gor， 在main这里是不会等待的
  //go readDoneWork(hasDoneWork, allDone)



  //time.Sleep(3 * time.Second)

  /**
  select {
  case <-hasDoneWork:
    fmt.Println("read done work...")
  }
  */

  ///**
  for {
    select {
    case <-hasDoneWork:
      workCount++
      fmt.Println("read done work...")
    case <-allDone:
      fmt.Println("all done 1...")
      return
    default:
      if workCount == 3 {
        fmt.Println("all done 2...")
        return
      }
    }
  }
  //*/

}
