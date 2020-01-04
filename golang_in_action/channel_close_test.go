package golang_in_action

import (
  "fmt"
  "sync"
  "testing"
  "time"
)

func dataProducer(ch chan int, wg *sync.WaitGroup)  {
  go func() {
    for i := 0; i < 10; i++ {
      ch <-i
      fmt.Println("生产了数据", i)
    }
    //close(ch)     // 这里不close， ok永远返回true

    time.Sleep(5 * time.Second)
    fmt.Println("sleep 5秒之后，再生产....")
    for i := 0; i < 5; i++ {
      ch <-i
      fmt.Println("生产了数据", i)
    }

    close(ch)     // 这里不close， ok永远返回true
    wg.Done()
  }()
}

func dataConsumer1(ch chan int, wg *sync.WaitGroup) {
  go func() {
    for i := 0; i < 11; i++ {     // 增加多读取， 已关闭的通道返回类型默认值
      data := <-ch
      fmt.Println("消费了数据", data)
    }
    wg.Done()
  }()
}

func dataConsumer2(ch chan int, wg *sync.WaitGroup) {
  go func() {
    for {
      if data, ok := <-ch; ok {
        fmt.Println("消费了数据", data)
      } else {
        //break       // 如果想一直循环消费， 可以注释掉
      }
    }
    wg.Done()
  }()
}


func TestComm(t *testing.T)  {
  var wg sync.WaitGroup
  ch := make(chan int)

  wg.Add(1)           // 添加一个协程运行生产者
  dataProducer(ch, &wg)

  wg.Add(1)           // 添加一个协程运行消费者
  //dataConsumer1(ch, &wg)
  dataConsumer2(ch, &wg)

  wg.Wait()
}

func TestCancel(t *testing.T) {
  cancelChan := make(chan struct{}, 0)      // 建立一个空的 channel， 表示为 cancel channel

  for i := 0; i < 10; i++ {
    go func(i int, cancelCh chan struct{}) {
      for {
        if isCanceled(cancelCh) {       // 一旦往 cancel channel写入数据，则表示要取消
          //fmt.Println("有一些不显示 // 2的输出，是因为协程还没有分配到执行这个i，没有sleep 5s")
          fmt.Println("已经cancel的输出-----------", i)
          //return
          break
        } else {
          fmt.Println("还没有cancel的输出, for一直在阻塞---------", i)
        }
        time.Sleep(5 * time.Second)
      }
      // 2
      //fmt.Println("8核的CPU， 默认8个协程同时跑---", i, "close channel之后，这里sleep 5秒都还会输出，证明close channel， 已经执行的gor还是会执行的, 不知道ctx cancel会不会也是这样？")
      //fmt.Println("还没有cancel的输出", i)
    }(i, cancelChan)
  }

  //fmt.Println("因为8核默认只跑8个协程，而且每个协程阻塞5秒， main协程阻塞3秒， 所以肯定少两个输出 //2 位置")
  //cancel2(cancelChan)
  cancel1(cancelChan)             // 广播机制，一旦close，则所有的 监听此channel的逻辑都直接返回
  time.Sleep( 3 * time.Second)
}









