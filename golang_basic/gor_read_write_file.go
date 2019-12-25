package main

import (
  "fmt"
  "io"
  "os"
  "sync"
)

type SyncWriter struct {
  m   sync.Mutex
  Writer io.Writer
}

func (w *SyncWriter) Write(b []byte) (n int, err error) {
  w.m.Lock()
  defer w.m.Unlock()
  return w.Writer.Write(b)
}

var data = []string {
  "hello",
  "aoh",
  "world",
}

func main() {
  f, err := os.Create("sample.txt")
  if err != nil {
    panic(err)
  }

  wr := &SyncWriter{sync.Mutex{}, f}
  wg := sync.WaitGroup{}

  for _, val := range data {
    wg.Add(1)
    go func(greeting string) {
      defer wg.Done()
      // fixme:  wr 这个参数是一个interface类型，所以传入的参数实体只要实现了 interface 里面的方法就可以了，这个就叫  "如果它走起步来像鸭子,并且叫声像鸭子, 那个它一定是一只鸭子。", Duck typing
      fmt.Fprintln(wr, greeting)
    }(val)
  }

  wg.Wait()
}