package main

import "testing"

func TestFillMatrix(t *testing.T) {
  //col := 3
  //row := 9
  var x [row][col]int
  fillMatrix(&x)
  t.Log(x)
}

