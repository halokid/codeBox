package main 

import (
  "fmt"
) 


/**
func main() {
  if (true) {
    defer fmt.Println("1")
  } else {
    defer fmt.Println("2")
  }

  fmt.Println("3")


  s := make([]int, 5, 10)
  s[0] = 100
  fmt.Println(s)
}
**/


/**
func main() {
  var a Integer = 1
  var b Integer = 2
  var i interface{} = &a

  sum := i.(*Integer).Add(b)
  fmt.Println(sum)
}



type Integer int 

func (a Integer) Add (b Integer) Integer {
  return a + b
}


func (a *Integer) Add(b Integer) Integer {
  return *a + b
}


type Fragment interface {
  Exec(transInfo *TransInfo) error
}


type GetPodAction struct {

}


func (g GetPodAction) Exec(transInfo *TransInfo) error {
  // ...

  return nil
}


**/


func main() {
  // i := 1
  c := complex(3, 12)
  fmt.Println(c)
}




































