package main
/**
两个 goroutine 会抢占性的输出
 */
import (
  "fmt"
  "runtime"
  "time"
)

var quit = make(chan int)

func loop() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d\n", i)
  }

  // 写入 quit channel
  // 这里是为了 如果 没有读取 quit channel 的话， 那么这个协程就会阻塞等待
  quit <- 0
}

func main() {
  t1 := time.Now()
  runtime.GOMAXPROCS(2)

  go loop()
  go loop()

  for i := 0; i < 2; i++ {
    <- quit
  }

  t2 := time.Now()
  fmt.Println(t2.Sub(t1))
}
