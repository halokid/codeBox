package main 

import (
  "fmt"
) 


// var n int = 99
// fmt.Println(n)

var (
  n int = 99
)
fmt.Println(n)


func main() {

  fmt.Println(n)
  /**
  // 这个是正常是赋值到类型， 不算是类型转换 
  type MyInt int 
  var i int = 1
  var j MyInt = 1
  **/

  

  /**
  // 错误的类型转换1
  type MyInt int 
  var i int = 1
  var j MyInt = (MyInt)i
  **/

  // /**
  // 正确的类型转换
  type MyInt int 
  var i int = 1
  var j MyInt = MyInt(i)
  // **/

  /**
  // 错误的类型转换2
  type MyInt int 
  var i int = 1
  var j MyInt = i.(MyInt)
  **/

  fmt.Println(i)
  fmt.Println(j)

  test()

  var m = 100
  fmt.Println(m)
}



func test() {
  var i int = 10

  var j = 5

  k := 8

  fmt.Println(i)
  fmt.Println(j)
  fmt.Println(k)
}






















