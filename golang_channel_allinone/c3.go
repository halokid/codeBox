package main

import (
  "fmt"
  "time"
)

func testClose() {
  sign := make(chan int, 2)

  ///**
  ch := make(chan int, 5)
  done := make(chan int)

  go func() {
    for i := 1; i < 5; i++ {
      ch <-i
      time.Sleep(time.Second)
    }

    close(ch)
    done <-1

    fmt.Println("the channel is closed")
  }()

  go func() {
    /**
    // 一直循环从 ch 拿数据, 根据ch取得数据的状态来判断退出
    for {
      i, ok := <-ch
      fmt.Printf("%d, %v\n", i, ok)

      if !ok {
        break
      }

      //time.Sleep(time.Second * 2)
      time.Sleep(time.Second)
    }
     */

    /**
    // 一直循环，以ch写入数据触发select， 判断读取ch的长度数量来判断退出
    i := 0
    for {
      if i >= 5 {
        break
      }
      select {
      case i := <-ch:
        fmt.Printf("for select get %d\n", i)
      }

      i++
    }
     */

    //i := 0
    for {
      //if i >= 10 {
      //  break
      //}
      select {
      case i := <-ch:
        fmt.Printf("select get %d\n", i)
      case <-done:
        fmt.Println("for done")
        //break       // 根本不能停止select监听
        return        // 可以停止select监听，跳出for循环
      }

      //i += 1
    }

  }()
  //*/

  <-sign
  <-sign
}

func main() {
  testClose()
}





