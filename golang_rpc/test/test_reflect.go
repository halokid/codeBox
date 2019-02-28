package test

import (
  "fmt"
  "reflect"
)

type Abc interface {
  String()  string
}

type Efg struct {
  data    string
}

func (e *Efg) String() string {
  return e.data
}

func GetEfg() *Efg {
  return nil
}

func CheckAE(a Abc) bool {
  return a == nil
}

func main() {
  efg := GetEfg()
  b := CheckAE(efg)
  fmt.Println(b)

  efgt := reflect.TypeOf(efg)
  fmt.Println(efgt)

  bt := reflect.TypeOf(b)
  fmt.Println(bt)

  //elem := bt.Elem()
  //fmt.Printf("%v\n", elem)
}

