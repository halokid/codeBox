package main 


import (
  "fmt"
  "reflect"
)

type MyStruct struct {
  name string
}


func (this *MyStruct) GetName() string {
  return this.name
}


func main() {
  a := new(MyStruct)
  a.name = "r00txx"
  
  typ := reflect.TypeOf(a)
  fmt.Println(typ)
  
  fmt.Println("--------------------")
  
  vpa := reflect.ValueOf(a)
  fmt.Println(vpa)
  
  fmt.Println("**********************************")
  
  var b MyStruct
  b.name = "abc"
  
  typ2 := reflect.TypeOf(b)
  fmt.Println(typ2)
  // fmt.Println(reflect.TypeOf(b)
  fmt.Println("--------------------")
  
  vpb := reflect.ValueOf(b)
  fmt.Println(vpb)
  
  fmt.Println("**********************************")
  
}








