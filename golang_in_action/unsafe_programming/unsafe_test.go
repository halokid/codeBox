package unsafe_programming

import (
  //"fmt"
  "sync"
  "sync/atomic"
  "testing"
  "time"
  "unsafe"
)

func TestUnsafe(t *testing.T) {
  i := 10
  //  1. unsafe.Pointer 取出指针
  // 2. 把指针强制转为 float64
  t.Log(&i)
  f := *(*float64)(unsafe.Pointer(&i))
  t.Log(unsafe.Pointer(&i))
  t.Log(f)
}


type MyInt int

func TestConvert(t *testing.T)  {
  a := []int{1, 2, 3, 4}
  b := *(*[]MyInt)(unsafe.Pointer(&a))
  t.Log(b)
}

func TestAtomic(t *testing.T) {
  var shareBufPtr unsafe.Pointer
  writeDataFn := func() {
    var data []int
    for i := 0; i < 10; i++ {
      data = append(data, i)
    }
    // fixme: 原子性操作，写入指针位置
    atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
  }

  readDataFn := func() {
    // fixme: 原来性操作，装载指针位置
    data := atomic.LoadPointer(&shareBufPtr)
    t.Log(data, *(*[]int)(data))
    //t.Log(data, *(*[]int)(data))
  }

  t.Log("------------------------------")
  var wg sync.WaitGroup
  writeDataFn()      // fixme: 首先为 shareBufPtr 创建内存位置，如果不跑这个逻辑，假如下面的读协程 比  写 协程更快的话， 那么就会出错， 先写是为了初始化 shareBufPtr 的内存位置， 因为var 仅仅只是生命，并没有初始化内存位置

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
      for i := 0; i < 10; i++ {
        writeDataFn()
        time.Sleep(100 * time.Microsecond)
      }
      wg.Done()
    }()

    wg.Add(1)
    go func() {
      for i := 0; i < 10; i++ {
        readDataFn()
        time.Sleep(100 * time.Microsecond)
      }
      wg.Done()
    }()
  }

  time.Sleep(1 * time.Second)
}














