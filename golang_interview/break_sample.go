package main 

import (
  "fmt"
) 

/** my note
控制语句

break 语句 经常用于中断当前 for 循环或跳出 switch 语句

Go 语言中 break 语句用于以下两方面：
1.用于循环语句中跳出循环，并开始执行循环之后的语句。
2.break在switch（开关语句）中在执行一条case后跳出语句的作用

**/


/**
func main() {
  var a int = 10

  for a < 20 {
    fmt.Printf(" a is : %d\n", a)
    a++ 

    if a > 15 {
      break
    }
  }
}
**/

// Break label 语句：我们在for多层嵌套时，有时候需要直接跳出所有嵌套循环， 这时候就可以用到go的label breaks特征了

var (
  a int
  b string
  c bool
)


var d, e int

// f := 10,  g := 20


func main() {
  fmt.Println("1")

Exit:
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      if i + j > 15 {
        fmt.Println("exit")
        break Exit 
      }
    }
  }

  fmt.Println("3")

  fmt.Println(a)
  fmt.Println(e)

  test()

  add(1, 2)
  add(1, 3, 7)
  add([]int{1, 6, 8}...)
}



func test() {
  xx := []int{1, 3, 7}
  fmt.Println(xx) 
}



func add(args ...int) int {
  sum := 0

  for _, arg := range args {
    sum += arg
  }

  fmt.Println(sum)
  return sum
}

















