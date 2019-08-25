package main

import "fmt"

func main() {
  c1 := make(chan interface{})
  close(c1)     // channel 虽然关闭了，但是还是可以继续其值的，但是不能写入了

  c2 := make(chan interface{})
  close(c2)
  fmt.Println("c1 & c2 defined done")

  var c1Count, c2Count int

  for i := 1000; i > 0; i-- {
    //fmt.Println("i ----------- ", i)
    select {
    case <-c1:
      c1Count++
      fmt.Println("c1 ------", <-c1)
    case <-c2:
      c2Count++
      fmt.Println("c2 ------", <-c2)
    }
  }

  // 尝试写入c1
  c1 <- 9       // 如果这里再尝试写入 c1， 这里会报错

  fmt.Println("c1Count: ", c1Count, "--- c2Count: ", c2Count)
}
