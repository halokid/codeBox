package main

import (
  "os"
  "testing"
)

func TestGreet(t *testing.T) {
  /*
  buffer := bytes.Buffer{}
  // todo: TDD， 早期我们的函数只是为了支持传入一个 bytes.Buffer 的指针，支持这个就可以了
  // todo: 后期假如我们的函数还需要支持的范围比 bytes.Buffer还要大，那怎么办呢？这里假如我们需要支持 os.Stdout
  //Greet(&buffer, "Chris")   // 早期
  got := buffer.String()
  want := "Hello, Chris"

  if got != want {
    t.Errorf("got '%s' want '%s'", got, want)
  }
  */

  Greet(os.Stdout, "Chris")   // 后期
}


