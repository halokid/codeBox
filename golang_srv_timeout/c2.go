package main

import (
  "context"
  "time"
)

func main() {
  d := time.Now().Add(50 * time.Millisecond)
  ctx, cancel := context.WithDeadline()
}
