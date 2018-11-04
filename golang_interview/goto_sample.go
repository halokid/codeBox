package main 

import (
  "fmt"
) 

/**
func main() {
  var i int 

  for {
    println(i)

    i++
    if i > 2 { goto BREAK }
  }

  BREAK:
    println("break")
  // EXIT:
    // println("exit")
}
**/


func main() {
  var a int = 10

  LOOP: 
  for a < 20 {
    if a == 15 {
      // 等于 15 的话， 则跳过迭代
      a = a + 1
      goto LOOP
    }

    fmt.Printf("a is : %d\n", a)
    a++
  }
}


/** my note
Golang支持在函数内 goto 跳转。goto语句与标签之间不能有变量声明。否则编译错误。
**/















