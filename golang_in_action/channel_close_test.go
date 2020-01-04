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
          break
        }
        time.Sleep(5 * time.Second)
      }
      fmt.Println(i, "Canceled")
    }(i, cancelChan)
  }

  //cancel2(cancelChan)
  cancel1(cancelChan)             // 广播机制，一旦close，则所有的 监听此channel的逻辑都直接返回
  time.Sleep( 3 * time.Second)
}









