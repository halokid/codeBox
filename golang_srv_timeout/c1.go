package main

import (
  "context"
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "sync"
  "time"
)

func calHandler(c *gin.Context)  {
  // 一个请求会触发调用三个服务， 每个服务输出一个 int
  // 请求要求结果为三个服务输出 int 之和
  // 请求返回时间不超过3秒， 大于3秒只输出已经获得的 int 之和
  var resContainer, sum int
  var success, resChan = make(chan int), make(chan int, 3)
  ctx, _ := context.WithTimeout(c, 3 * time.Second)

  go func() {
    for {
      select {
      case resContainer = <- resChan:
        sum += resContainer
        fmt.Println("add ", resContainer)
      case <- success:
        fmt.Println("result: ", sum)
        return
      case <- ctx.Done():
        fmt.Println("result by ctx.Done() : ", sum)
      }
    }
  }()

  wg := sync.WaitGroup{}
  wg.Add(3)

  go func() {
    resChan <- microService1()
  }()

  go func() {
    resChan <- microService2()
  }()

  go func() {
    resChan <- microService3()
  }()

  wg.Wait()
  success <- 1

  return
}

func main() {
  r := gin.New()
  r.GET("/calxx", calHandler)
  http.ListenAndServe(":8080", r)
}

func microService1() int {
  time.Sleep(1 * time.Second)
  return 1
}

func microService2() int {
  time.Sleep(2 * time.Second)
  return 2
}

func microService3() int {
  time.Sleep(10 * time.Second)      // 超时
  return 3
}

























