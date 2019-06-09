package main
/**
这个是在运用多路复用写法之前的比较蠢的方法， 也能实现多路复用的效果
 */
import (
  "fmt"
  "time"
)

func sendxx(ch chan int, gap time.Duration) {
  i := 0
  for {
    i++
    ch <- i
    time.Sleep(gap)
  }
}

func collect(source chan int, target chan int) {
  for v := range source {
    target <- v
  }
}

func recvxx(ch chan int) {
  for v := range ch {
    fmt.Printf("recv %d\n", v)
  }
}

func main() {
  var ch1 = make(chan int)
  var ch2 = make(chan int)
  var ch3 = make(chan int)

  go sendxx(ch1, time.Second)
  go sendxx(ch2, 2 * time.Second)

  go collect(ch1, ch3)
  go collect(ch2, ch3)

  recvxx(ch3)
}

















