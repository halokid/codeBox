package golang_in_action

import (
  "fmt"
  "sync"
  "testing"
)

func TestSyncPool(t *testing.T) {
  pool := &sync.Pool{
    New: func() interface{} {
      fmt.Println("创建一个新对象")
      return 100
    },
  }

  v := pool.Get().(int)
  fmt.Println(v)

  pool.Put(3)
  //runtime.GC()	//GC 会清除sync.pool中缓存的对象
  v1, _ := pool.Get().(int)
  fmt.Println(v1)

  v2, _ := pool.Get().(int)
  fmt.Println(v2)

}

func TestSyncPoolMulti(t *testing.T) {
  pool := &sync.Pool{
    New: func() interface{} {
      fmt.Println("创建一个新对象")
      return 10
    },
  }

  pool.Put(100)
  pool.Put(100)
  pool.Put(100)

  var wg sync.WaitGroup
  for i := 0; i < 10; i++ {
    go func(i int) {
      wg.Add(1)
      fmt.Println(pool.Get())
      wg.Done()
    }(i)
  }
  wg.Wait()
}






