package main

import (
  "log"
  "time"
)

func solution(digits string, num string) int {
  start := time.Now()
  for _, n := range num {
    for _, d := range digits {
      if n == d {
        log.Printf("%+v", n)
        continue
      }
    }
  }
  cost := time.Since(start)
  log.Println(cost)
  log.Println(int64(cost / time.Microsecond))   // todo: time durarion to int64
  return 0
}

func main() {
  solution("1234567890", "012")
}
