package main

import (
  "fmt"
  "math/rand"
  "time"
)

/**
开启多个线程做赚钱和花钱的操作，共享读写remainMoney这个剩余金额变量，实现生产者消费者模型
同步控制模型，生产者模型
 */

var lockChan = make(chan int, 1)
var remainMoney = 1000

func testSynchronize()  {
  quit := make(chan bool, 2)

  go func() {
    for i := 0; i < 10; i++ {
      money := (rand.Intn(12) + 1) * 100
      go testSynchronize_expense(money)

      time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
    }
    quit <-true
  }()

  go func() {
    for i := 0; i < 10; i++ {
      money := (rand.Intn(12) + 1) * 100
      go testSynchronize_gain(money)

      time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
    }
    quit <-true
  }()

  <-quit
  <-quit
  fmt.Println("主协程退出")
}

func testSynchronize_expense(money int) {
  lockChan <-0      // 锁住lockChan, 阻塞，目的有两个：  1. 不会执行下面的逻辑  2. lockChan <-0写进不了，达到原子操作的作用

  if remainMoney >= money {
    srcRemainMoney := remainMoney
    remainMoney -= money
    fmt.Printf("原来有%d, 花了%d, 剩余%d\n", srcRemainMoney, money, remainMoney)
  } else {
    fmt.Printf("想消费%d钱不够了，只剩余%d\n", money, remainMoney)
  }

  <-lockChan        // 释放
}

func testSynchronize_gain(money int) {
  lockChan <-0

  srcRemainMoney := remainMoney
  remainMoney += money
  fmt.Printf("原来有%d, 赚了%d, 剩余%d\n", srcRemainMoney, money, remainMoney)

  <-lockChan
}

func main() {
  testSynchronize()
}






