package main 

import (
  "fmt"
  "sync"
  "net/http"
)


func main() {
  var wg sync.WaitGroup
  var urls = []string {
          "http://www.baidu.com/",
          "http://www.qq.com",
          "http://www.163.com",
}

for _, url := range urls {
  //increate the waitgroup counter.
  wg.Add(1)
  //launch a goroutine to fetch the URL.
  go func(url string) {
    //decrement the counter when the goroutine complete 
    defer wg.Done()
    //fetch the URL
    http.Get(url)
    
    fmt.Println(url)
  }(url)
}
//wait for all HTTP fetched to complete
wg.Wait()
fmt.Println("over")
}