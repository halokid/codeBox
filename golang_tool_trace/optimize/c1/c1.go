package main

import (
  "fmt"
  "os"
  "runtime"
  "runtime/trace"
  "sync"
  "sync/atomic"
  "time"
)

var (
  stop  int32
  count int64
  sum   time.Duration
)

func concat1() {
  // 一次过创建gor
  for n := 0; n < 100; n++ {
    for i := 0; i < 8; i++ {
      go func() {
        s := "GO GC"
        s += " " + "Hello"
        s += " " + "World"
        _ = s
      }()
    }
  }
}

func concat2() {
  // 分批创建gor
  wg := sync.WaitGroup{}
  for n := 0; n < 100; n++ {
    wg.Add(8)
    for i := 0; i < 8; i++ {
      go func() {
        s := "GO GC"
        s += " " + "Hello"
        s += " " + "World"
        _ = s

        wg.Done()
      }()
    }
    wg.Wait()
  }
}

func main() {
  f, _ := os.Create("trace.out")
  defer f.Close()
  trace.Start(f)
  defer trace.Stop()

  go func() {
    var t time.Time
    for atomic.LoadInt32(&stop) == 0 {
      t = time.Now()
      runtime.GC()      // 计算GC的时间， 不是显式调用GC
      sum += time.Since(t)
      count++
    }
    fmt.Printf("GC平均花费时间: %v\n", time.Duration(int64(sum) / count))
  }()

  //concat1()
  concat2()
  // todo: 一旦stroe值1给stop之后, atomic.LoadInt32(&stop) == 0  条件就不会命中
  atomic.StoreInt32(&stop, 1)
}







