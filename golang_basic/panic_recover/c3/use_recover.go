package main

import (
  "log"
  "time"
)

type Work int

func server(workChan chan Work)  {
  for work := range workChan {
    go safelyDo(work)
  }
}

func safelyDo(work Work) {
  defer func() {
    // todo: recover() 会返回runtime时的error
    if err := recover(); err != nil {
      log.Printf("Work失败 %s, %v, 可以进行recover的操作", err, work)
    }
  }()

  do(work)
}

func do(work Work) error {
  if work == 5 {
    s := "i等于5的时候失败"
    panic(s)      // todo: panic才能引发recover
    //err := errors.New(s)    // todo: 不能引发recover
    //return err
  }
  log.Printf("正常工作 ----- %+v", work)
  return nil
}

func main() {
  workChan := make(chan Work, 10)
  var w Work
  //w = 9
  for w = 9; w > 0; w-- {
    log.Printf("w ---------- %+v", w)
    workChan <-w
  }
  server(workChan)

  time.Sleep(5 * time.Second)
}



