package test

import (
  "testing"
  "crypto/md5"
  "fmt"
)

func Test_getMd5(t *testing.T) {
  str := "abc123"
  data := []byte(str)
  mds := md5.Sum(data)
  md5str := fmt.Sprintf("%x", mds)
  fmt.Println(md5str)
}
