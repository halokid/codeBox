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


type Small struct {
  a int
}

var pool = sync.Pool{
  New: func() interface{} {
    return new(Small)
  },
}

func inc(s *Small) {
  s.a++
}

func BenchmarkWithoutPool(b *testing.B) {
  var s *Small
  for i := 0; i < b.N; i++ {
    for j := 0; j < 50; j++ {
      s = &Small{ 1 }
      b.StopTimer(); inc(s); b.StartTimer()
    }
  }
}

func BenchmarkWithPool(b *testing.B) {
  var s *Small
  for i := 0; i < b.N; i++ {
    for j := 0; j < 50; j++ {
      s = pool.Get().(*Small)
      s.a = 1
      b.StopTimer(); inc(s); b.StartTimer()
      pool.Put(s)
    }
  }
}

















