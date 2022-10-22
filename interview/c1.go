package main

import "strings"

//import "flag"

func main() {
  //var addr = flag.String("addr", "127.0.0.1:8080", "The addr of the application.")
}

func solution(A, B string) bool {
  tmp := ""
  for c := range A {
    tmp += string(c)
  }

  for d := range B {
    if !strings.Contains(string(d), tmp) {
      return false
    }
  }
  return true
}



