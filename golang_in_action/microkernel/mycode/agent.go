package mycode

import (
  "context"
  "errors"
  "fmt"
  "strings"
  "sync"
)

type State int

const (
  Running   State = iota
  Waiting
)

var WrongStateError = errors.New("agent不能正常响应, 状态异常")

// Collector
type CollectorsError struct {
  CollectErrors  []error
}

func (ce CollectorsError) Error() string {
  var strs []string
  for _, err := range ce.CollectErrors {
    strs = append(strs, err.Error())
  }
  return strings.Join(strs, ";")
}


// Event
type Event struct {
  Source      string
  Content     string
}

// ------------------ 定义接口 --------------------
type EventReciver interface {
  OnEvent(evt Event)
}


type Collector interface {
  Init(evtReceiver EventReciver) error
  Start(agtCtx context.Context) error
  Stop() error
  Destory() error
}


type Agent struct {
  collectors    map[string]Collector
  evtBuf        chan Event
  cancel        context.CancelFunc
  ctx           context.Context
  state         State
}

func (agt *Agent) EventProcessGroutine() {
  var evtSeg  [10]Event
  for {
    for i := 0; i < 10; i++ {
      select {
      case evtSeg[i] = <-agt.evtBuf:
      case <-agt.ctx.Done():
        return

      }
    }
    fmt.Println(evtSeg)
  }
}

func NewAgent(sizeEvtBuf int) *Agent {
  agt := Agent{
    collectors:     map[string]Collector{},
    evtBuf:         make(chan Event, sizeEvtBuf),
    state:          Waiting,
  }
  return &agt
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
  if agt.state != Waiting {
    return WrongStateError
  }
  agt.collectors[name] = collector
  return collector.Init(agt)
}

func (agt *Agent) startCollectors() error {
  var err error
  var errs CollectorsError
  var mutex sync.Mutex
  for name, collector := range agt.collectors {

    go func(name string, collector Collector, ctx context.Context) {
      defer func() {
        mutex.Unlock()
      }()

      // fixme: 定义接口的一个好处就是，写代码结构的时候，不用考虑接口的具体实现，继承接口的定义写下去
      err = collector.Start(ctx)
      mutex.Lock()
      if err != nil {
        errs.CollectErrors = append(errs.CollectErrors, errors.New("start发生错误: " + name + ":" + err.Error()))
      }
    }(name, collector, agt.ctx)

  }   // END FOR

  return errs
}

func (agt *Agent) stopCollectors() error {
  var err error
  var errs CollectorsError
  for name, collector := range agt.collectors {
    if err = collector.Stop(); err != nil {
      errs.CollectErrors = append(errs.CollectErrors, errors.New("stop发生错误: " + name + ":" + err.Error()))
    }
  }
  return errs
}

func (agt *Agent) destoryCollectors() error {
  var err error
  var errs CollectorsError
  for name, collector := range agt.collectors {
    if err = collector.Destory(); err != nil {
      errs.CollectErrors = append(errs.CollectErrors, errors.New("destory发生错误: " + name + ":" + err.Error()))
    }
  }
  return errs
}


func (agt *Agent) Start() error {
  if agt.state != Waiting {
    return WrongStateError
  }
  agt.state = Running
  // fixme: 定义了 cancel函数
  agt.ctx, agt.cancel = context.WithCancel(context.Background())
  go agt.EventProcessGroutine()     // 仅仅是为了每收到10个事件就输出一次
  return agt.startCollectors()
}

func (agt *Agent) Stop() error {
  if agt.state != Waiting {
    return WrongStateError
  }
  agt.state = Waiting
  // fixme: 调用了cancle函数
  agt.cancel()
  return agt.stopCollectors()
}

func (agt *Agent) Destory() error {
  if agt.state != Waiting {
    return WrongStateError
  }
  return agt.destoryCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
  agt.evtBuf <-evt
}





