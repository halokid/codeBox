package hello

import (
  "fmt"
  "testing"
)

func TestHello(t *testing.T) {
  got := hello()
  expect := "hello func in package hello." 

  if got != expect {
    t.Error("got [%s] expected [%s]", got, expect)
  }
}


func BenchmarkHello(b *testing.B) {
  for i := 0; i < b.N; i++ {
    hello()
  } 
}


func ExampleHello() {
  hl := hello()
  fmt.Println(hl)
}


/**
func TestStackGrowth(t *testing.T) {
  t.Parallel()
  var wg sync.WaitGroup

  wg.Add(1)
  go func() {
    defer wg.Done()
    growStack()
  }()

  wg.Wait()
}


func BenchemarkValueRead (b *testing.B) {
  var v Value
  v.Store(new(int))
  b.RunParalle(func(pb *testing.PB)) {
    for pb.Next() {
      x := v.Load().(*int)
      for *x != 0 {
        b.Fatal("wrong value: got %v, want 0", *x)
      }
    } 
  }
}

**/