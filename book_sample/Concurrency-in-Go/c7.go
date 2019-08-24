package main
/**
sync cond 主要用于在各个channel中实现状态通知的
 */

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  c := sync.NewCond(&sync.Mutex{})
  queue := make([]interface{}, 0, 10)

  removeFromQueue := func(delay time.Duration) {
    time.Sleep(delay)
    c.L.Lock()
    queue = queue[1:]     // 把第一个元素抛出去
    fmt.Println("remoed from queue")
    c.L.Unlock()
    c.Signal()
  }

  for i := 0; i < 10; i++ {
    c.L.Lock()
    for len(queue) == 2 {
      // 如果数组是 2 位的长度，则停止处理
      // fixme: 在 wait 之前，必须要先lock起来，没有lock 是 wait 不了的，必须锁住才能等待啊
      c.Wait()
    }
    fmt.Println("adding to queue")
    //c.L.Lock()      // 写在这里会产生错误
    queue = append(queue, struct {}{})
    go removeFromQueue(1 * time.Second)
    c.L.Unlock()
  }
}






