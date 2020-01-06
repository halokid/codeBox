package testing

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSquare(t *testing.T) {
  inputs := [...]int{1, 2, 3}
  expected := [...]int{1, 4, 9}

  for i := 0; i < len(inputs); i++ {
    ret := square(inputs[i])
    if ret != expected[i] {
      t.Errorf("输入 %d, 输出  %d, 实际正确结果为 %d", inputs[i], expected[i], ret)
    }
  }
}

func TestErrorInCode(t *testing.T)  {
  fmt.Println("Start")
  t.Error("Error")
  fmt.Println("End")
}

func TestFatalInCode(t *testing.T)  {
  fmt.Println("Start")
  t.Fatal("Fatal")
  fmt.Println("End")
}

func TestSquareAssert(t *testing.T)  {
  inputs := [...]int{1,2,3}
  expected := [...]int{1,4,9}
  for i := 0; i < len(inputs); i++ {
    ret := square(inputs[i])
    assert.Equal(t, expected[i], ret)
    //fmt.Println(ret, expected[i])
  }
}

