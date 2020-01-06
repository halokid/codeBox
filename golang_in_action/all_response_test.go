package golang_in_action

import (
  "fmt"
  "runtime"
  "testing"
  "time"
)

func doTask(id int) string {
  time.Sleep(10 * time.Second)
  return fmt.Sprintf("the result is from %d", id)
}

func AllResponse() string {
  numOfRunner := 10
  ch := make(chan string, numOfRunner)
  for i := 0; i < numOfRunner; i++ {
    go func(i int) {
      ret := doTask(i)
      ch <-ret
    }(i)
  }

  finalRet := ""
  for j := 0; j < numOfRunner; j++ {
    finalRet += <-ch + "\n"
  }
  return finalRet
}

func TestAllResponse(t *testing.T) {
  t.Log("Before:", runtime.NumGoroutine())
  t.Log(AllResponse())
  time.Sleep(1 * time.Second)
  t.Log("After:", runtime.NumGoroutine())
}








