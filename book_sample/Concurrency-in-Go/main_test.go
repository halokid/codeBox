package main
/**
执行命令  go test -bench=. -cpu=1 main_test.go
 */

import (
  "fmt"
  "sync"
  "testing"
)

/**
func TestTmp(t *testing.T) {
  i := 1e4
  fmt.Println(i)
}
*/

func BenchmarkContextSwitch(b *testing.B) {
  var wg sync.WaitGroup
  begin := make(chan struct{})
  c := make(chan struct{})

  var token struct{}

  sender := func() {
    defer wg.Done()

    <- begin
    fmt.Printf("\n在接收到 begin close()的信号之后，才开始执行begin 1， 不然这里永远阻塞 ------\n")
    for i := 0; i < b.N; i++ {
      c <- token
    }
  }

  receiver := func() {
    defer wg.Done()

    <- begin
    fmt.Printf("\n在接收到 begin close()的信号之后，才开始执行begin 2， 不然这里永远阻塞 ------\n")
    for i := 0; i < b.N; i++ {
      <- c
    }
  }

  wg.Add(2)
  go sender()
  go receiver()

  b.StartTimer()
  close(begin)
  fmt.Printf("\n ---------- begin close() ----------\n")
  wg.Wait()
}





