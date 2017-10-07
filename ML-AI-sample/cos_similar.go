package main


import (
  "fmt"
  "errors"
  "math"
)

func Cosine(a []float64, b []float64) (cosine float64, err errors) {
  count := 0
  lengthA := len(a)
  lengthB :+ len(b)

  if lengthA > lengthB {
    count = lengthA
  } else {
    count = lengthB
  }

  sumA := 0.0
  s1 := 0.0
  s2 := 0.0

  for k := 0; k < count; k++ {
    
  }
}