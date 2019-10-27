package main

import (
  "errors"
  "fmt"
)

func GetFibonacci(n int) ([]int, error) {
  if n < 2 || n > 100 {
    return nil, errors.New("n should be in [2, 100]")
  }

  fibSl := []int{1, 1}
  for i := 2; i < n; i++ {
    fibSl = append(fibSl, fibSl[i-2] + fibSl[i-1])
  }
  return fibSl, nil
}

func main() {
  if v, err := GetFibonacci(10); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(v)
  }
}
