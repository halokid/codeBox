package main

import "time"

func main() {
  doWork := func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
    heartbeat := make(chan interface{})
    results := make(chan time.Time)

    go func() {
      defer close(heartbeat)
      defer close(results)

      pulse := time.Tick(pulseInterval)
      workGen := time.Tick(2 * pulseInterval)

      sendPulse := func() {
        select {
        case heartbeat <- struct {}{}:
        default:

        }
      }

      sendResult := func() {
        for {
          select {
          case <-done:
            return
          case <-pulse:
            sendPulse()
          case results <-r:
            return
          }
        }
      }

      for {
        select {
        case <-done:
          return
        case <-pulse:
          sendPulse()
        case r := <-workGen:
          sendResult()
        }
      }
    }()
    return heartbeat, results
  }
}




