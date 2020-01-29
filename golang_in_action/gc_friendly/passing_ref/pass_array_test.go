package passing_ref

import "testing"

/**
运行查看 GC 情况
GODEBUG=gctrace=1  go test -bench=BenchmarkPassingArrayWithValue

如果是用 指针引用的话， 很明显 GC 的次数 比 内存引用  的次数要少很多， 所以指针引用对GC更友好，所以性能更高
 */
const NumOfElems = 1000

type Content struct {
  Detail  [10000]int
}

func withValue(arr [NumOfElems]Content) int {
  // 内存引用
  return 0
}

func withReference(arr *[NumOfElems]Content) int {
  // 指针引用
  return 0
}

func TestFn(t *testing.T) {
  var arr [NumOfElems]Content
  withValue(arr)
  withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
  var arr [NumOfElems]Content

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    withValue(arr)
  }
  b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
  var arr [NumOfElems]Content

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    withReference(&arr)
  }
  b.StopTimer()
}





