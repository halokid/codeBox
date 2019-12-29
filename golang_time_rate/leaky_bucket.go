package main

/**
漏桶算法
当桶里的水满的时候，则拒绝服务
每服务一次，则添加水
 */
import (
  "fmt"
  "math"
  "sync"
  "time"
)

type LeakyBucket struct {
  rate       float64 //固定每秒出水速率
  capacity   float64 //桶的容量
  water      float64 //桶中当前水量
  lastLeakMs int64   //桶上次漏水时间戳 ms
  lock       sync.Mutex
}

func (l *LeakyBucket) Allow() bool {
  l.lock.Lock()
  defer l.lock.Unlock()

  // 执行漏水
  now := time.Now().UnixNano() / 1e6        // 纳秒转为微秒
  // eclipse就是水龙头的遮挡， 流出越多，则water越少， 流出越少，则water越多
  // fixme: rate 越大， 则流出越多
  eclipse := float64(now - l.lastLeakMs) * l.rate / 1000      // 每一次处理请求漏出来的水量
  fmt.Println("eclipse ----------- ", eclipse)
  l.water = l.water - eclipse
  fmt.Println("l.water ----------- ", l.water)
  l.water = math.Max(0, l.water)
  l.lastLeakMs = now
  if (l.water + 1)  < l.capacity {
    // 水没有加满
    l.water++
    return true
  } else {
    return false
  }

}

func (l *LeakyBucket) Set(r, c float64) {
  l.rate = r
  l.capacity = c
  l.water = 0
  l.lastLeakMs = time.Now().UnixNano() / 1e6
}

func main() {
  var wg sync.WaitGroup
  var lr LeakyBucket
  lr.Set(3, 3)

  for i := 0; i < 10; i++ {
    wg.Add(1)

    fmt.Println("发起请求", i, time.Now())
    go func(i int) {
      defer wg.Done()
      if lr.Allow() {
        fmt.Println("返回请求", i, time.Now())
      } else {
        fmt.Println("请求", i, "超过漏桶速度，请求被拒绝....")
      }
    }(i)

    time.Sleep(100 * time.Millisecond)
  }
  wg.Wait()
}





