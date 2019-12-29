package main

import (
  "fmt"
  "sync"
  "time"
)

type LimitRate struct {
  rate  int           //计数周期内最多允许的请求数
  begin time.Time     //计数开始时间
  cycle time.Duration //计数周期
  count int           //计数周期内累计收到的请求数
  lock  sync.Mutex
}

func (l *LimitRate) AllowRefuse() bool {
  // 超过速率直接被拒绝
  l.lock.Lock()
  defer l.lock.Unlock()

  if l.count == l.rate - 1 {
    now := time.Now()
    // 如果间隔的时间是在 周期之外
    // fixme: 当一开始发起请求时，因为第一个1秒肯定不是从 0秒开始计算的， 所以第一个1秒只能请求两次
    if now.Sub(l.begin) >= l.cycle {
      l.Reset(now)
      return true
    } else {
      return false
    }
  } else {
    l.count++
    return true
  }
}

func (l *LimitRate) AllowWait() bool {
  // 超过速率阻塞等待，直到可以达到继续请求的条件
  l.lock.Lock()
  defer l.lock.Unlock()

  if l.count == l.rate - 1 {
    for {     // 阻塞
      now := time.Now()
      if now.Sub(l.begin) >= l.cycle {
        l.Reset(now)
        return true
      } else {
        // fixme: 可发现 每一秒做多有3个 "返回请求"
        // fmt.Println("请求超过速率，等待中...........")
      }
    }
  } else {
    l.count++
    return true
  }
}

func (l *LimitRate) Set(r int, cycle time.Duration) {
  l.rate = r
  l.begin = time.Now()
  l.cycle = cycle
  l.count = 0
}

func (l *LimitRate) Reset(t time.Time) {
  l.begin = t
  l.count = 0
}

func main() {
  var wg sync.WaitGroup
  var lr LimitRate
  //lr.Set(3, 2 * time.Second)    // fixme: 两秒最多处理 3 个请求
  lr.Set(3, time.Second)    // fixme: 一秒最多处理 3 个请求

  for i := 0; i < 100; i++ {
    wg.Add(1)
    fmt.Println("发起请求:", i, time.Now())
    go func(i int) {
      defer wg.Done()
      //if lr.AllowRefuse() {
      if lr.AllowWait() {
        fmt.Println("返回请求:", i, time.Now())
      } else {
        fmt.Println("请求", i, "超过计数器速度，请求被拒绝....")
      }
    }(i)

    time.Sleep(200 * time.Millisecond)
  }
  wg.Wait()
}












