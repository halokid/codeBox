package main

import "fmt"

/**
bridge channel模式
 */

func main() {
  bridge := func(done <-chan interface{},
                  chanStream <-chan <-chan interface{}) <-chan interface{} {
    valStream := make(chan interface{})     // 1
    go func() {
      defer close(valStream)

      for {
        var stream <-chan interface{}
        select {
        case maybeStream, ok := <-chanStream:
          if ok == false {
            return
          }
          stream = maybeStream
        case <-done:
          return
        }

        for val := range orDone(done, stream) {
          select {
          case valStream <-stream:
            fmt.Println(val)
          case <-done:
          }
        }

      }
    }()
    return valStream
  }
  fmt.Println(bridge)
}

func orDone(done <-chan interface{}, stream <-chan interface{}) string {
  //fmt.Println(done)
  return ""
}
