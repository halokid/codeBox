package main

import (
  "github.com/spf13/cast"
  "testing"
)

func TestBytesShow(t *testing.T) {
  s := "abc"
  t.Log("[]byte(s) --------- ", []byte(s))

  x := 64
  t.Log("[]byte(x) --------- ", []byte(cast.ToString(x)))
}
