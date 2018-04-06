package main

import (
  "fmt"
)

/**
关于这个并行模型， 整个流程可以说是非常操蛋的。。。
1. 首先建立两个 chan 类型的变量，  然后把这两个变量赋予 函数 Processor， 从 Processor(origin, wait)，下面就是 赋值到  origin channel 的流程,
  Processor(origin, wait) 这一行是不会立马执行的， 因为参数是 channel 类型的， 所以要等 channel全部 读完才执行， 直到 close(origin)，

  <-wait 这个很重要， 这个是表示程序一直要等 wait 这个 channel 消费完才能结束


2、  再去到  Processor 函数， 这个就是一个 递归函数， 这个递归函数其实每一轮结束的标志就是 这一句 out <- num


整个程序总结就是很抽象。。。。。。。。。。。。

一开始 prime的值其实是2来的， 传进去函数之后， 清晰点的理解， 关键可 out 的值，  out每当2， 3， 5， 7，  9 的时候才有值，
所以为了 各个 数都是并行去匹配, 所以才把 out 也设置成 channel

所以整个 并行模型 的总结就是， 为了我们需要的逻辑都实现并行， 就把相关处理函数的 参数 设置为 channel， 然后利用递归函数， 不断
把参数一直 赋值 给函数， 取得函数的返回作为 我们需要的结果就可以了
 */

func main() {
  origin, wait := make(chan int), make(chan int)
  Processor(origin, wait)
  for num := 2; num < 10; num++ {
    origin <- num   //把num写进 chan 去,  其实这里没有任何并行的概念， 真正的并行逻辑是从 go func(){} 开始的
  }
  close(origin)
  <-wait
}

func Processor(seq chan int, wait chan int) {
  go func() {
    prime, ok := <-seq
    //fmt.Println("get prime -------------------")
    //fmt.Println(prime)
    //fmt.Println("++++++++++++++\n\n\n")
    if !ok {
      close(wait)
      return
    }
    //fmt.Println(prime)
    out := make(chan int)
    Processor(out, wait)     //因为参数是 channel类型的， 所以这里是不会马上执行的, 而是去到了下一个 for 循环
    //其实这个程序还是不够好的， 因为当5 碰到 不整除的时候
    for num := range seq {
      //fmt.Println("get num -------------------")
      //fmt.Println(num)
      fmt.Println("get prime -------------------")
      fmt.Println(prime)
      fmt.Println("\n\n\n")

      //fmt.Println("\n\n")
      if num % prime != 0 {     //操， 是我想得太复杂了， 其实这个 prime 的值一直都是 2
        //fmt.Println(num)
        out <- num
      }
    }
    close(out)
  }()
}



