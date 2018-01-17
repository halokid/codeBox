/*
 * @Author: r00x
 * @Date: 2017-10-07 23:03:23 
 * @Last Modified by: mikey.zhaopeng
 * @Last Modified time: 2017-10-07 23:04:25
 */


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
    if k >= lengthA {
      s2 += math.Pow(b[k], 2)
      continue
    }

    if k >= lengthB {
      s1 += math.Pow(a[k], 2)
    }

    sumA += a[k] * b[k]
    
    s1 += math.Pow(a[k], 2)
    s2 += math.Pow(b[k], 2)
  }

  if s1 == 0 || s2 == 0 {
    return 0.0, errors.New("vectors should not br null(all zero)")
  }

  // 下面的算法就是：  两个二维数组的值的 乘积的累计相加  除以  两个二维数组的每个值平方之和的平方根求值
  return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

func main() {
  cos, err := Cosine([]float64{0, 1, 0, 1, 1}, []float64{1, 0, 1, 0, 0})
  if err != nil {
    panic(err)
  }

  fmt.Println(cos)
}



