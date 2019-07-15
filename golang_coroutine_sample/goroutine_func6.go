package main

import (
  "fmt"
  "github.com/mozillazg/request"
  "github.com/r00tjimmy/ColorfulRabbit"
  "net/http"
  "sync"
  "time"
)

func GetWhyy(wg *sync.WaitGroup) {
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



func main() {
  t1 := time.Now()

  // 下面的这种写法，并不能达到协程的最大利用， 耗时一共 36s
  for i := 0; i < 10; i++ {
    var wg sync.WaitGroup

    for j := 0; j < 20; j++ {
      wg.Add(1)
      GetWhyy(&wg)
    }

    wg.Wait()
  }

  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}













