package main

import "testing"

const url = "http:xxxx.com"

func TestAdd(t *testing.T) {
  s := Add(url)
  if s == "" {
    t.Errorf("Test.Add error")
  }
}

func BenchmarkAdd(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Add(url)
  }
}