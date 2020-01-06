package bdd

import (
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

func TestSpec(t *testing.T) {
  Convey("给出两个偶数", t, func() {
    a := 2
    b := 4

    Convey("当两个偶数相加", func() {
      c := a + b

      Convey("那么结果肯定也是偶数", func() {
        So(c % 2, ShouldEqual, 0)
      })
    })
  })
}
