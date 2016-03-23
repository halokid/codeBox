package main 

import (
  "fmt"
)

func main() {
  var a int = 20
  var ip *int
  
  ip = &a
  
  fmt.Printf("memory address of val a:  %x\n", &a)
  
  fmt.Printf("memory address of point ip:   %x\n", ip)
  
  fmt.Printf("the val *ip value is: %d\n", *ip)
}