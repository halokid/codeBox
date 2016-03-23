package main 

import (
  "fmt"
)

func main() {
  var a int = 100
  var b int = 200
  
  fmt.Printf("before change, value a is:  %d\n", a)
  fmt.Printf("before change, value b is:  %d\n", b)
  
  swap(&a, &b)
  
  fmt.Printf("after change, value a is: %d\n", a)
  fmt.Printf("after change, value b is: %d\n", b)
}

func swap(x *int, y *int) {
  var temp int
  temp = *x 
  *x = *y
  *y = temp
}