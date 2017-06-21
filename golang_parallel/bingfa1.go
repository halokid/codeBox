package main

import (
  "sync"
  "fmt"
  //"time"
  "runtime"
)

var count int = 0;

func counter(lock *sync.Mutex) {
  lock.Lock()
  count++
  fmt.Println(count)
  lock.Unlock()
}


func main() {
  lock := &sync.Mutex{}
  for i := 0; i < 10; i++ {
    go counter(lock)
  }
  //time.Sleep(10 * time.Second)

  for {
    lock.Lock()     //这两个锁和解锁的意义不大吧??因为有没有输出都是一样的
    c := count      // 必须等这个赋值完之后，才能解锁
    lock.Unlock()

    runtime.Gosched()   //wait for goroutine process finish
    if c == 10 {
      fmt.Println("goroutine end")
      break
    }
  }
}







