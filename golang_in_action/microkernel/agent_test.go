package microkernel

import (
"context"
"errors"
"fmt"
"testing"
"time"
)

type DemoCollector struct {
  evtReceiver	EventReceiver
  agtCtx		context.Context
  stopChan	chan struct{}
  name		string
  content		string
}

func NewCollect(name string, content string) *DemoCollector {
  return &DemoCollector{
    stopChan:	make(chan struct{}),
    name:		name,
    content:	content,
  }
}

func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
  fmt.Println("initialize collect", c.name)
  c.evtReceiver = evtReceiver
  return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
  fmt.Println("start collect", c.name)
  for {
    select {
    case <-agtCtx.Done():
      c.stopChan <- struct{}{}
      break
    default:
      // 默认是每50秒注册一个事件
      time.Sleep(time.Microsecond * 50)
      c.evtReceiver.OnEvent(Event{c.name, c.content})
    }
  }
}

func (c *DemoCollector) Stop() error {
  fmt.Println("stop collect", c.name)
  select {
  case <-c.stopChan:
    return nil
  case <-time.After(time.Second * 1):
    return errors.New("failed to stop for timeout")
  }
}

func (c *DemoCollector) Destory() error {
  fmt.Println(c.name, "released resources.")
  return nil
}

/**
[{c2 2} {c2 2} {c1 1} {c2 2} {c1 1} {c1 1} {c2 2} {c1 1} {c2 2} {c1 1}]
[{c2 2} {c2 2} {c1 1} {c2 2} {c1 1} {c1 1} {c2 2} {c1 1} {c2 2} {c1 1}]
[{c2 2} {c2 2} {c1 1} {c1 1} {c2 2} {c1 1} {c2 2} {c1 1} {c2 2} {c1 1}]
[{c2 2} {c1 1} {c2 2} {c1 1} {c2 2} {c2 2} {c1 1} {c2 2} {c1 1} {c1 1}]
[{c2 2} {c1 1} {c2 2} {c2 2} {c1 1} {c2 2} {c1 1} {c2 2} {c1 1} {c1 1}]

刚好输入 5 行， 则是 50个， 每个并行等待 time.Microsecond * 50
2500 / 50 = 50
 */

func TestAgent(t *testing.T) {
  agt := NewAgent(100)
  c1 := NewCollect("c1", "1")
  c2 := NewCollect("c2", "2")
  agt.RegisterCollector("c1", c1)
  agt.RegisterCollector("c2", c2)
  agt.Start()
  //fmt.Println(agt.Start())
  //time.Sleep(time.Second * 1)
  time.Sleep(time.Microsecond * 2500)
  // fixme: cancel当前的 context， 那么由当前 context所产生的所有协程都会cancel
  agt.Stop()
  agt.Destory()
}






