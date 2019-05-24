package main

import (
  "log"
  "net/http"
  _ "net/http/pprof"
)

func main() {
  go func() {
    for {
      log.Println(Add("xxxxx"))
    }
  }()

  http.ListenAndServe("0.0.0.0:6060", nil)
}
