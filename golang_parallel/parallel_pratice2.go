package main

import (
  "fmt"
  "math/rand"
)


/**
func rand_generator_1() int {
  return rand.Int()
}
**/


func rand_generator_2() chan int {
  out := make(chan int)
  
  go func() {
    for {
      out <- rand.Int()
      // fmt.Println("1111")
    }
  }()
  // fmt.Println()
  return out
}


func rand_generator_3() chan int {
    // 创建两个随机数生成器服务
    rand_generator_1 := rand_generator_2()
    rand_generator_2 := rand_generator_2()

    //创建通道
    out := make(chan int)

    //创建协程
    go func() {
        for {
            //读取生成器1中的数据，整合
            out <- <-rand_generator_1
        }
    }()
    go func() {
        for {
            //读取生成器2中的数据，整合
            out <- <-rand_generator_2
        }
    }()
    return out
}



func main() {
  rand_service_handler := rand_generator_3()
  fmt.Printf("%d\n", <-rand_service_handler)
}


































