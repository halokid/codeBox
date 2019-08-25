package main 

import (
  "fmt"
  // "runtime"
  // "sync"
) 


/**
func main() {
  runtime.GOMAXPROCS(1)
  wg := sync.WaitGroup{} 
  wg.Add(20)

  for i := 0; i < 10; i++ {
    go func() {
      fmt.Println("A: ", i)
      wg.Done()
    }()
  }


  for i := 0; i < 10; i++ {
    go func(i int) {
      fmt.Println("B: ", i)
      wg.Done()
    }(i)
  }

  wg.Wait() }

**/

/**
// golang 的OOP继承
type People struct {}

func (p *People) showA() {
  fmt.Println("showA")
  p.showB()
}

func (p *People) showB() {
  fmt.Println("showB") 
}


type Teacher struct {
  People
}

func (t *Teacher) showB() {
  fmt.Println("teacher showB") 
}

func main() {
  t := Teacher{}
  t.showA()
  t.showB()
}
**/


/**

考点：select随机性
解答：
select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：

select 中只要有一个case能return，则立刻执行。
当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
如果没有一个case能return则可以执行”default”块。

**/
/**
func main() {
  runtime.GOMAXPROCS(1)

  intChan := make(chan int, 1)
  stringChan := make(chan string, 1)

  intChan <- 1
  stringChan <- "hello"

  select {
    case value := <- intChan:
      fmt.Println(value)
    case value := <- stringChan:
      panic(value)
  }   
}

**/




/**
func calc(index string, a, b int) int {
  ret := a + b 
  fmt.Println(index, a, b, ret)
  return ret
}


func main() {
  a := 1
  b := 2

  defer calc("1", a, calc("10", a, b)) 
  a = 0

  defer calc("2", a, calc("20", a, b))
}
**/

func main() {
  s := make([]int, 5)
  s = append(s, 1, 2, 3)
  fmt.Println(s)

  x := make([]int, 0)
  x = append(x, 1, 2, 3)
  fmt.Println(x)
}





































