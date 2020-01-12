package mycode

import (
  "context"
  "fmt"
  "errors"
  "testing"
  "time"
)

type DemoCollector struct {
  evtReceiver     EventReciver
  agtCtx          context.Context
  stopChan        chan struct{}
  name            string
  content         string
}

func NewCollect(name string, content string) *DemoCollector {
  return &DemoCollector{
    stopChan:       make(chan struct{}),
    name:           name,
    content:        content,
  }
}

func (c *DemoCollector) Init(evtReceiver EventReciver) error {
  fmt.Println("初始化collect", c.name)
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
    return errors.New("stop超时失败")
  }
}

func (c *DemoCollector) Destory() error {
  fmt.Println(c.name, "Destory资源")
  return nil
}

func TestAgent(t *testing.T) {
  agt := NewAgent(10)
  c1 := NewCollect("c1", "1")
  c2 := NewCollect("c2", "2")

  agt.RegisterCollector("c1", c1)
  agt.RegisterCollector("c2", c2)

  //agt.Start()
  fmt.Println(agt.Start())
  time.Sleep(time.Second * 1)
  agt.Stop()
  agt.Destory()
}
























