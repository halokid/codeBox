package main

import (
  "log"
  "net/http"
  "time"
)


func Logger(inner http.Handler, name string) http.Handler {
// 因为 handlers.go 这句 "Handler(route.handler)"， 所以返回 http.Handler

  return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        inner.ServeHTTP(w, r)     //依然进行 ServeHTTP 服务，然后下面的逻辑就是装饰 HTTP 的log展示了
        
        log.Printf(
          "%s\t%s\t%s\t%s",
          r.Method,
          r.RequestURI,
          name,
          time.Since(start),
        )
  } )
}

