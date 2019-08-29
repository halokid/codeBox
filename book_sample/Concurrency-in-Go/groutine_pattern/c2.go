package main

import "fmt"
/**
这种范式是比较好点的， 有一个专门的逻辑对channel进行写， 一个专门进行读，这样不会污染数据，在并行范式中污染 gor 之间的数据，会造成很恐怖的后果
 */

func main() {

  chanOwner := func() <-chan int {
    results := make(chan int, 5)      // 作用域在这里， 限制了只能在这里写入channel

    go func() {
      defer close(results)
      for i := 0; i <= 5; i++ {
        results <-i
      }
    }()

    return results
  }

  consumer := func(results <-chan int) {
    for result := range results {
      fmt.Println("readed: ", result)
    }
    fmt.Println("done readed..")
  }

  results := chanOwner()        // 限制了只能读取
  consumer(results)
}



