package main

import (
  "fmt"
  "net/http"
)

func main() {
  checkStatus := func(done <-chan interface{}, urls ...string,) <-chan *http.Response {
    responses := make(chan *http.Response)

    go func() {
      defer close(responses)

      for _, url := range urls {
        resp, err := http.Get(url)
        if err != nil {
          fmt.Println("err -----", err)
          continue
        }

        // 这种方式就没办法结束 gor
        responses <-resp

        /**
        // 用下面这种方式，就是假如有done数据传过来，可以结束该gor
        select {
        case <-done:
          return
        case responses <-resp:
        }
        */
      }
    }()

    return responses
  }

  done := make(chan interface{})
  defer close(done)       // 其实注释掉这里也可以，但是不是一种好的方式

  urls := []string{"http://www.baidu.com", "http://badhost"}
  for response := range checkStatus(done, urls...) {
    fmt.Println("resp: ", response.Status)
  }
}









