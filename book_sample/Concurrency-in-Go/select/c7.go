package main

import (
  "fmt"
)

func doChan(c chan int) {
  //defer close(c)
  c <-1
  //i++
  //close(c)
}

func main() {
  c := make(chan int)
  //defer close(c)

  //doChan(c)     // 这样写必错，因为channel 是要在不同的channel通信的， 必须用go起不同的gor
  i := 0

  go doChan(c)

  go doChan(c)

  //select {
  //case <-c:
  // fmt.Println("c readed..")
  //}

  //time.Sleep(2 * time.Second)

  ///**
  for {
   select {
   case <-c:
     fmt.Println("for 不断监听c")
     i++
      //break
     //return
   //default:     // 这样写就是sb，因为如果程序还来不及对c写入的话，就默认就走这里， 那就是还来不及写入就已经close了channel，还搞毛啊
   //  close(c)
   default:
     if i == 2{
       return
       //break
     } else {
       fmt.Println("run .....")
     }
   }
  }
  //*/

  /**
 go func() {
   for {
   select {
   case <-c:
     fmt.Println("for 不断监听c")
      //break
     //return
   //default:     // 这样写就是sb，因为如果程序还来不及对c写入的话，就默认就走这里， 那就是还来不及写入就已经close了channel，还搞毛啊
   //  close(c)
   }
  }
 }()
  */

  //close(c)

  //time.Sleep(2 * time.Second)
  fmt.Println("如果不加for，select读取到一个 c channel的写入，马上返回，其他的channel对 c 的写入读取不了")
}
