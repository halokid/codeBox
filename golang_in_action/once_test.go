package golang_in_action

import (
  "fmt"
  "sync"
  "testing"
  "unsafe"
)

type Singleton struct {

}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
  once.Do(func() {
    fmt.Println("创建一个 Singleton obj")
    singleInstance = new(Singleton)
    //singleInstance = make(Singleton)
  })
  return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
  var wg sync.WaitGroup
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
      obj := GetSingletonObj()      // 协程抢占式创建实例
      fmt.Printf("%x\n", unsafe.Pointer(obj))
      wg.Done()
    }()
  }
  wg.Wait()
}
