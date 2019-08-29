package main

import "fmt"

/**
  在主  gor 定义了 data， 然后在某一个 gor 去访问这个 data, 下面这样的写法可能造成data的数据污染
 */


func main() {
  data := make([]int, 4)

  loopData := func(handleData chan<- int) {
    defer close(handleData)
    // 这里是拿外部的 data 变量
    for i := range data {
      handleData <-data[i]
    }
  }

  handleData := make(chan int)
  // 这里又定义了一个 gor， 也是用到外部的data数据
  go loopData(handleData)

  for num := range handleData {
    fmt.Println(num)
  }
}



