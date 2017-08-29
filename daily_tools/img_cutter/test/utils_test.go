package main

import (
  "testing"
  "../utils"
)

func Test_utils(t *testing.T) {
  var  err  error
  utils.CheckErr("testing", err)
}
