package main

import (
  "fmt"
  "sync"
)

func main() {
  fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
    var wg sync.WaitGroup
    multiplexedStream := make(chan interface{})

    multiplex := func(c <-chan interface{}) {
      defer wg.Done()
      for i := range c {
        select {
        case <-done:
          return
        case multiplexedStream <-i:

        }
      }
    }

    wg.Add(len(channels))
    for _, c := range channels {
      go multiplex(c)
    }

    go func() {
      wg.Wait()
      close(multiplexedStream)
    }()

    return multiplexedStream
  }

  //fmt.Println(fanIn())

}
