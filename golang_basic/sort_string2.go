package main 

import (
  "fmt"
  "sort"
  "strings"
)

func SortString(w string) string {
  s := strings.Split(w, "")
  sort.Strings(s)
  return strings.Join(s, "")
}


func main() {
  w1 := "bcad"
  w2 := SortString(w1)
  
  fmt.Println(w1)
  fmt.Println(w2)
}

