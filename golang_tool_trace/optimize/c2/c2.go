package main

import (
  "fmt"
  "net/http"
  _ "net/http/pprof"
  "sync"
)

func newBuf() []byte {
  return make([]byte, 10<<20)
}

// 使用 sync.Pool 复用需要的buf
var bufPool = sync.Pool{
  New: func() interface{} {
    return make([]byte, 10 << 20)
  },
}


func main() {
  go func() {
    http.ListenAndServe(":6060", nil)
  }()

  http.HandleFunc("/example2", func(w http.ResponseWriter, r *http.Request) {
    //b := newBuf()
    b := bufPool.Get().([]byte)

    // 模拟执行一些工作
    for idx := range b {
      b[idx] = 1
    }

    fmt.Fprintf(w, "done, %v", r.URL.Path[1:])
  })

  http.ListenAndServe(":8080", nil)
}
