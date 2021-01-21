package main

import (
  "sync"
  "testing"
)

func TestSyncPool(t *testing.T) {
  p := &sync.Pool{
    New: func() interface{} {
      return 0
    },
  }

  a := p.Get().(int)
  p.Put(1)
  b := p.Get().(int)

  t.Log("a: ", a, ", b:", b)
}
