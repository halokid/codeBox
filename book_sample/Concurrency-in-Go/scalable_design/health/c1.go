package main

import (
  "log"
  "time"
)

func main() {
  type startGoroutineFn func(done <-chan interface{},
    pulseInterval time.Duration) (heartbeat <-chan interface{}) //1

  newSteward := func(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn { //2

    return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {

      heartbeat := make(chan interface{})

      go func() {
        defer close(heartbeat)

        var wardDone chan interface{}
        var wardHeartbeat <-chan interface{}
        startWard := func() { //3

          wardDone = make(chan interface{}) //4
          wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2) //5
        }
        startWard()
        pulse := time.Tick(pulseInterval)

      monitorLoop:

        for { //6
          timeoutSignal := time.After(timeout)

          for {

            select {
            case <-pulse:

              select {
              case heartbeat <- struct{}{}:

              default:

              }
            case <-wardHeartbeat: //7
              continue monitorLoop

            case <-timeoutSignal: //8

              log.Println("steward: ward unhealthy; restarting")
              close(wardDone)
              startWard()
              continue monitorLoop
            case <-done:

              return
            }

          }

        }
      }()

      return heartbeat

   }

  }
}
