package golang_in_action

import (
  "fmt"
  "testing"
  "time"
)

func TestObjPool(t *testing.T) {
  pool := NewObjPool(10)
  for i := 0; i < 11; i++ {
    if v, err := pool.GetObj(time.Second * 1); err != nil {
      t.Error(err)
    } else {
      fmt.Printf("%T\n", v)

      // fixme: 假如不执行 release的逻辑，那么11个的话，则会超出池的容量，最后一个获取会超时
      if err := pool.Release(v); err != nil {
       t.Error(err)
      }
    }
  }
  fmt.Println("Done")
}
