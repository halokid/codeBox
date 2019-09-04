package main
/**
超时和取消
 */

func main() {
  var value interface{}
  done := make(chan int)
  valueStream := make(chan int)
  select {
  case <-done:
    return
  case value = <-valueStream:

  }

  result := reallyLongCalculation(value)

  select {
  case <-done:
    return
  case valueStream <-result:

  }

  reallyLongCalcultion := func(done <-chan interface{}, value interface{}) interface{} {

  }
}
