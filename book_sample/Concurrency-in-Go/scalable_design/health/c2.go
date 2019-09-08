package main

import (
  "log"
  "os"
  "time"
)

func main() {
  log.SetOutput(os.Stdout)
  log.SetFlags(log.Ltime | log.LUTC)

  doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
    log.Println("ward: Hello, i am irresponsible")

    go func() {
      <-done
      log.Println("ward: i am halting")
    }()

    return nil
  }

}















