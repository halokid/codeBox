package main
/**
总结：  channel 要在不同的 groutine，或者内存空间 中去进行通信， 不然会造成死锁
 */

import (
  "fmt"
  "time"
)

func writeCh(c chan int) {
  c <-1
}

func readCh(c chan int) {
  select {
  case <-c:
    fmt.Println("c channel 被读取")
  }
}

func main() {
  //var c chan int
  //c = make(chan int)

  c := make(chan int)

  // fixme: 假如在这里定义了这句，则会产生错误, fatal error: all goroutines are asleep - deadlock! , goroutine 1 [chan send]:
  // 因为send 一个 1 给 c 之后， 这个 main 的 groutine 就阻塞住了，除非有其他的 groutine 去读取（消费）这个 c channel，不然的话，就会死锁
  //c <-1

  // 封装了另外一个 func 来写入（生产）c channel， 所以加上这句不会死锁
  go writeCh(c)
  //close(c)


  //<-c

  /**
  // 下面这个逻辑是在 main 这个 groutine 去读取（消费）c channel 的
  select {
  case <-c:
    fmt.Println("c read..")
  }
  */

  // 或者封装另外一个 func（也就是 groutine, 只要不是同一个内存空间就可以了)， 所以下面两种写法都可以
  //readCh(c)
  go readCh(c)

  fmt.Println("如果在 main 里面定义了 c <-1， 那么就不会运行这里")

  //close(c)

  fmt.Println(c)
  time.Sleep(2 * time.Second)
}
