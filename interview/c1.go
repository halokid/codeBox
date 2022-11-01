package main

import "log"

//import "flag"

type G struct {
  a string
}
var g G

func main() {
  //var addr = flag.String("addr", "127.0.0.1:8080", "The addr of the application.")
  //x :=  interface{}(&G{"a"})
  //y :=  interface{}(&G{"a"})
  //log.Println(x == y)
  var s1 []int
  s2 := []int{1, 2, 3}
  n1 := copy(s1, s2)
  log.Println("--", n1)
  log.Println("--", s1, s2)
  log.Println("--", s1 == nil)
}

//func solution(A, B string) bool {
//  tmp := ""
//  for c := range A {
//    tmp += string(c)
//  }
//
//  for d := range B {
//    if !strings.Contains(string(d), tmp) {
//      return false
//    }
//  }
//  return true
//}



