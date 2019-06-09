package main
/**
这个就是 多路复用的典型例子， 是优化了c10的写法
 */
import (
  "fmt"
  "time"
)

func sendyy(ch chan int, gap time.Duration) {
  i := 0
  for {
    i++
    ch <- i
    time.Sleep(gap)
  }
}

func recvyy(ch1 chan int, ch2 chan int) {
  for {
    select {
    case v := <- ch1:
      fmt.Printf("recv %d from ch1\n", v)
    case v := <- ch2:
      fmt.Printf("recv %d from ch2\n", v)
    }
  }
}

func checkSend(ch1 chan int, ch2 chan int) {
  v := 0
  for {
    select {
    case ch1 <- v:
      fmt.Printf("get send from ch1")
    case ch2 <- v:
      fmt.Printf("get send from ch2")
    }
  }
}


func main() {
  var ch1 = make(chan int)
  var ch2 = make(chan int)

  go sendyy(ch1, time.Second)
  go sendyy(ch2, 2 * time.Second)

  recvyy(ch1, ch2)
}









