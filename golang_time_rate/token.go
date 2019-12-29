package main

import (
  "fmt"
  "sync"
  "time"
)

type TokenBucket struct {
  rate         int64 //固定的token放入速率, r/s
  capacity     int64 //桶的容量
  tokens       int64 //桶中当前token数量
  lastTokenSec int64 //桶上次放token的时间戳 s

  lock sync.Mutex
}

func (l *TokenBucket) Allow() bool {
  l.lock.Lock()
  defer l.lock.Unlock()

  now := time.Now().Unix()
  l.tokens = l.tokens + (now - l.lastTokenSec) * l.rate       // 先添加令牌
  if l.tokens > l.capacity {
    l.tokens = l.capacity
  }
  l.lastTokenSec = now
  if l.tokens > 0 {
    l.tokens--
    return true
  } else {
    return false
  }
}

func (l *TokenBucket) Set(r, c int64) {
  l.rate = r
  l.capacity = c
  l.tokens = 0
  l.lastTokenSec = time.Now().Unix()
}

func main() {
  var wg sync.WaitGroup
  var lr TokenBucket
  //lr.Set(3, 3)   // 每秒添加3个令牌， 则表达每秒速率限制为3个请求， 桶的容量为3
  lr.Set(3, 90)   // fixme: 这里等待了30秒， 添加了90个令牌， 如果一开始桶里面的令牌足够多，也可以处理超出速率限制的请求，因为本来桶里就有很多令牌， 当桶里的这些令牌用完的时候， 才是完全按照从 0 个令牌开始，遵守速率限制来处理请求

  time.Sleep(30 * time.Second)   // 主程序阻塞 1s 以便让桶中储备好 3 个令牌
  for i := 0; i < 10; i++ {
    wg.Add(1)

    fmt.Println("发起请求", i, time.Now())
    go func(i int) {
      defer wg.Done()
      if lr.Allow() {
        fmt.Println("返回请求", i, time.Now())
      } else {
        fmt.Println("超过速率， 拒绝请求", i)
      }
    }(i)
    time.Sleep(200 * time.Millisecond)    // 每 1s 创建 5 个请求, 因为这里sleep 0.2秒
  }
  wg.Wait()
}








