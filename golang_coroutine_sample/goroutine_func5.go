package main

import (
  "fmt"
  "github.com/mozillazg/request"
  "github.com/r00tjimmy/ColorfulRabbit"
  "net/http"
  "sync"
  "time"
)

func GetWh(wg *sync.WaitGroup) {
  defer wg.Done()

  url := "https://www.tianqiapi.com/api/?version=v1&city=%E6%B5%8E%E5%8D%97"
  c := new(http.Client)
  req := request.NewRequest(c)
  rsp, err := req.Get(url)
  if err != nil {
    fmt.Println("err get url")
  }
  fmt.Println(rsp.StatusCode)
}

func GetWhx() {
  //defer wg.Done()

  url := "https://www.tianqiapi.com/api/?version=v1&city=%E6%B5%8E%E5%8D%97"
  c := new(http.Client)
  req := request.NewRequest(c)
  rsp, err := req.Get(url)
  if err != nil {
    fmt.Println("err get url")
  }
  fmt.Println(rsp.StatusCode)
}

func main() {
  var wg sync.WaitGroup

  t4 := time.Now()
  GetWhx()        // 单个的请求接口， 时间大概用250ms
  t5 := time.Now()

  t6 := ColorfulRabbit.DurTime(t4, t5)
  fmt.Println(t6)
  println("----------------------------")

  // --------------------------------------------------
  t1 := time.Now()

  // 200 url请求, 采用协程的方式, 大概性能能压缩到，经测试大概时间只需要 按顺序执行， 15个 url 请求的时间
  for i := 0; i < 10; i++ {
    for j := 0; j < 20; j++ {

      // 这里用协程， 时间大概是  3.5s
      wg.Add(1)
      go GetWh(&wg)

      // 这里不用协程， 时间大概是  42s
      //GetWhx()
    }
  }

  wg.Wait()

  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}

















