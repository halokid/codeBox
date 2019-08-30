package main

import (
  "fmt"
  "time"
)

func writeChan(c chan int) error {
  c <-1
  return nil
}

func main() {
  c := make(chan int)
  go writeChan(c)

  for {
    // fixme: 如果加上这句， 则是等待 c channel 的写入， 那么肯定是走到  case <-c
    //time.Sleep(2 * time.Second)
    select {
    case <-c:
      fmt.Println("c readed")
      return
      //break
    default:
     time.Sleep(2 * time.Second)
     fmt.Println("for select default..")
      //break
      return
    }
  }

  // 因为整个 main 都return 了， 所以绝对不会运行到这里
  fmt.Println("break for loop...")
}
