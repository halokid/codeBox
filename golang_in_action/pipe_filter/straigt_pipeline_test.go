package pipe_filter

import "testing"

func TestNewStraightPipeline(t *testing.T) {
  spliter := NewSpliterFilter(",")
  converter := NewToIntFilter()
  sum := NewSumFilter()

  sp := NewStraightPipeline("p1", spliter, converter, sum)
  ret, err := sp.Process("1,2,3")
  if err != nil {
    t.Fatal(err)
  }
  if ret != 6 {
    t.Fatal("期望值是6， 但是计算的实际值是", ret)
  }
  t.Log("pipeline-filter架构模式计算正确", ret)
}
