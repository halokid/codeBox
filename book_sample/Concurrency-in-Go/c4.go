package main
/**
  当闭包的循环，去添加 goroutine 的时候， goroutine 里面的闭关变量， 是会抢占式输出slice里面的元素的
  只是执行得太快， 所以 输出是  "good day"
 */
import (
  "fmt"
  "sync"
)

func main() {
  var wg3 sync.WaitGroup

  for _, salutation := range []string{"hello", "greetings", "good day"} {
  //for i := 0; i < 10; i++  {
    wg3.Add(1)

    go func() {
      defer wg3.Done()
      fmt.Println(salutation)       // 抢占式输出slice的内容
      //fmt.Println(i)    // 输出i之后，表示for里面添加 goroutine， 并不是顺序输出的
    }()
  }

  wg3.Wait()


  var wg4 sync.WaitGroup

  for _, salutation := range []string{"hello", "greetings", "good day"} {
    wg4.Add(1)

    // salutation的副本传递给这个函数， 也就是 salutation 的副本传递给闭包， 闭包会 holding 住这个内存的赋值，一直到func结束
    // 下面的 goroutine 还是会抢占式声明函数的， 所以上面三个元素是乱序输出
    go func(salutation string) {
      defer wg4.Done()
      fmt.Println(salutation)
    }(salutation)
  }

  wg4.Wait()

}














