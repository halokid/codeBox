package golang_in_action

import (
  "fmt"
  "testing"
  "time"
)

func service() string {
  time.Sleep(50 * time.Millisecond)
  return "Done 3333"
}

func otherTask() {
  fmt.Println("working on someting else 2222")
  time.Sleep(100 * time.Millisecond)
  fmt.Println("Task is done 2222")
}

func TestService(t *testing.T) {
  fmt.Println(service())
  otherTask()
}

func AsyncService() chan string {
  retCh := make(chan string, 1)
  //retCh := make(chan string)
  go func() {
    ret := service()
    fmt.Println("returned result 1111")
    retCh <-ret             // 1
    fmt.Println("如果是缓冲channel， 则这里不会阻塞到最后才执行 service exited 1111")
  }()
  return retCh
}

func TestAsyncService(t *testing.T) {
  retCh := AsyncService()
  otherTask()
  fmt.Println(<-retCh)          // 2
}




