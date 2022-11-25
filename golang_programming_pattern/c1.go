package main

import (
  "log"
  "strings"
)

// TODO: map1 function
func MapStrToStr(arr []string, fn func(s string) string) []string {
  var newArray []string
  for _, it := range arr {
    newArray = append(newArray, fn(it))
  }
  return newArray
}

// TODO: map2 function
func MapStrToInt(arr []string, fn func(s string) int) []int {
  var newArray []int
  for _, it := range arr {
    newArray = append(newArray, fn(it))
  }
  return newArray
}

// ----------------------------------------------------------
// TODO: reduce
func Reduce(arr []string, fn func(s string) int) int {
  sum := 0
  for _, it := range arr {
    sum += fn(it)
  }
  return sum
}

// ----------------------------------------------------------
// TODO: Filter
func Filter(arr []int, fn func(n int) bool) []int {
  var newArray []int
  for _, it := range arr {
    if fn(it) {
      newArray = append(newArray, it)
    }
  }
  return newArray
}

func main() {
  var list = []string{"hellox", "world", "all"}

  x := MapStrToStr(list, func(s string) string {
    return strings.ToUpper(s)
  })
  log.Printf("%v\n", x)

  y := MapStrToInt(list, func(s string) int {
    return len(s)
  })
  log.Printf("%v\n", y)

  // ----------------------------------------------------------
  z := Reduce(list, func(s string) int {
    return len(s)
  })
  log.Printf("%v", z)

  // ----------------------------------------------------------
  var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  out := Filter(intset, func(n int) bool {
    return n % 2 == 1
  })
  log.Printf("%v", out)
}










