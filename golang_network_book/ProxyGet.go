package main

import (

)
import (
  "os"
  "fmt"
  "net/url"
  "net/http"
  "net/http/httputil"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Println("usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
    os.Exit(1)
  }

  proxyString := os.Args[1]
  proxyURl, err := url.Parse(proxyString)
  checkError(1, err)

  rawURL := os.Args[2]
  url, err := url.Parse(rawURL)
  checkError(2, err)

  transport := &http.Transport{Proxy:http.ProxyURL(proxyURl)}
  client := &http.Client{Transport:transport}

  request, err := http.NewRequest("GET", url.String(), nil)

  dump, _ := httputil.DumpRequest(request, false)
  fmt.Println(string(dump))

  response, err := client.Do(request)

  checkError(3, err)
  fmt.Println("read ok")

  if response.Status != "200 OK" {
    fmt.Println(response.Status)
    os.Exit(2)
  }

  fmt.Println("response OK")

  ///**
  // 为什么这里要用 for 循环来读取 HTML 包的大小， 而不是用一个固定的值来读取呢？？？
  //  因为你又不知道你访问的 HTML 包具体的大小是多少， 所以你设置固定的值的话就会产生程序的错误
  // 所以就用这种方法， 先设置一个固定的大小值，然后 for 循环去按照这个固定的大小值来读取 HTML 包
  var buf [512]byte
  reader := response.Body
  for {
    n, err := reader.Read(buf[0:])
    if err != nil {
      os.Exit(0)
    }
    fmt.Println(string(buf[0:n]))
    fmt.Println("-------------------------------------------")
    fmt.Println(n)
  }
    //**/

  /**
    var buf [20000]byte
    reader := response.Body
    reader.Read(buf[0:])
    fmt.Println(string(buf[0:]))
    fmt.Println("print buf")
  **/

  os.Exit(0)

}


func checkError(code int, err error) {
  if err != nil {
    fmt.Println("Error")
    os.Exit(code)
  }
}





















