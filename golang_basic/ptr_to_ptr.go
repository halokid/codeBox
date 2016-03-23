package main 

import (
  "fmt"
)

func main() {
  var a int
  var ptr *int
  var pptr **int
  
  a = 3000
  
  ptr = &a  //point to a, address of a 
  pptr = &ptr   //point to ptr, address of ptr
  
  fmt.Printf("val a = %d\n", a)
  fmt.Printf("the value of the point *ptr = %d\n", *ptr)
  fmt.Printf("the value of the point to point **ptr = %d\n", **pptr)
}