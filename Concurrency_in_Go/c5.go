package main
/**
 对 channel 进行非阻塞读
 */
import (
  "fmt"
  "time"
)

func main() {
  myChan := make(chan string)

  go func() {
    myChan <- "Message"
  }()

  select {
  case msg := <- myChan:
    fmt.Println(msg)
  default:
    fmt.Println("No Msg")
  }

  <-time.After( time.Second * 2)

  select {
  case msg := <- myChan:
    fmt.Println(msg)
  default:
    fmt.Println("No Msg")
  }

  // 对 channle 进行非阻塞写
  select {
  case myChan <- "message":
    fmt.Println("sent the message")
  default:
    fmt.Println("no message sent")
  }

  <-time.After(time.Second * 2)

  go func() {
    m := <- myChan
    fmt.Println(m)
  }()

  <-time.After(time.Second * 2)

}




