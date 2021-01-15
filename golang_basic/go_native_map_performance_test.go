package main

import "testing"

func test(m map[int]int) {
  for i := 0; i < 10000; i++ {
    m[i] = i
  }
}

func BenchmarkMap(b *testing.B) {
  for i := 0; i < b.N; i++ {
    b.StopTimer()
    m := make(map[int]int)
    b.StartTimer()
    test(m)
  }
}

func BenchmarkMapCap(b *testing.B) {
  for i := 0; i < b.N; i++ {
    b.StopTimer()
    m := make(map[int]int, 10000)
    b.StartTimer()
    test(m)
  }
}

