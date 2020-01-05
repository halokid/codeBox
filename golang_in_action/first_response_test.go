package golang_in_action

import (
  "fmt"
  "runtime"
  "testing"
  "time"
)

func runTask(id int) string {
  time.Sleep(10 * time.Second)
  return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
  numOfRunner := 10
  //ch := make(chan string)

  // fixme: 起了一个有长度的 buffer channel， 这样假如一旦写入了 ch <-ret， 这样一旦有其中一个 gor跑完了 runTask函数（也就是说函数有返回）， 那么就会马上触发   return <-ch（原来是一直阻塞的），则整个 FirstResponse函数都会返回， 则会退出 gor 的运行时， 所有的 gor都会退出
  ch := make(chan string, numOfRunner)
  for i := 0; i < numOfRunner; i++ {
    fmt.Println("runtime gor num:", runtime.NumGoroutine())
    go func(i int) {
      ret := runTask(i)
      ch <-ret
    }(i)
  }
  return <-ch
}

func TestFirstResponse(t *testing.T) {
  // fixme: 用了 buffer channel之后， 这里会输出2 ，所以证明必要一个channel 返回， 其他的gor都会退出，2是最少的协程数量， 因为main func一个，  返回的gor一个， 一共2个
  t.Log("Before:", runtime.NumGoroutine())
  t.Log(FirstResponse())
  time.Sleep(1 * time.Second)
  t.Log("After:", runtime.NumGoroutine())
  t.Log("before  和  after 的协程数是一样的， 证明没有协程泄漏, 也就是说没有往close的 channel写东西， 也没有协程在做无意义的等待")
}
