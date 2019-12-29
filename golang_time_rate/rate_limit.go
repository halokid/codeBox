package main

import (
  "golang.org/x/time/rate"
  "net/http"
)

var limiter = rate.NewLimiter(2, 5)   // 速率为2， 桶的容量为5

func limit(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if !limiter.Allow() {
      http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
      return
    }
    next.ServeHTTP(w, r)
  })
}

func okHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK\n"))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", okHandler)
  http.ListenAndServe(":4000", limit(mux))
}



