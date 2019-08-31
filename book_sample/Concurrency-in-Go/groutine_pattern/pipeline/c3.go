package main

import (
  "fmt"
  "math/rand"
)

func main() {

  repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
    valueStream := make(chan interface{})
    go func() {
      defer close(valueStream)
      for {
        select {
        case <-done:
          return
        case valueStream <-fn():

        }
      }
    }()
    return valueStream
  }

  take := func(done <-chan interface{}, valueStream <-chan interface{}, num int, ) <-chan interface{} {

    takeStream := make(chan interface{})
    go func() {
      defer close(takeStream)
      for i := 0; i < num; i++ {
        select {
        case <-done:
          return
        case takeStream <- <-valueStream:
        }
      }
    }()
    return takeStream
  }

  done := make(chan interface{})
  defer close(done)

  randx := func() interface{} {
    return rand.Int()
  }

  fmt.Println("randx ------- ", randx)

  for num := range take(done, repeatFn(done, randx), 10) {
    fmt.Println("num ----- ", num)
  }
}











