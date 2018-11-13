package main 

import (
  "fmt"
  "runtime"
  "sync"
) 

/**
如果map由多协程同时读和写就会出现 fatal error:concurrent map read and map write的错误

如下代码很容易就出现map并发读写问题
**/

// ERROR CODE 1
type UserAges struct {
  ages map[string]int
  sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
  ua.Lock()
  defer ua.Unlock()

  ua.ages[name] = age
}


func (ua *UserAges) Get(name string) int {
  if age, ok := ua.ages[name]; ok {
    return age
  }
  return -1
}



// ERROR CODE 2
func main() {
  c := make(map[string]int)

  // 开一个协程写map
  go func() {
    for j := 0; j < 1000; j++ {
      c[fmt.Sprintf("%d", j)] = j
    }
  }()

  //  开一个协程读map 
  go func() {
    for j := 0; j < 1000; j++ {
      fmt.println(c[fmt.Sprintf("%d", j)]) 
    }
  }

  time.Sleep(time.Second * 20)
}



/**
解决方法
**/

// 1. 通用锁
type Demo struct {
  Data map[string]string
  Lock sync.Mutex
}


func (d Demo) Get(k string) string {
  d.Lock.Lock()

  defer d.Lock.Unlock()

  return d.Data[k]
}


func (d Demo) Set(k, v string) {
  d.Lock.Lock()

  defer d.Lock.Unlock()

  d.Data[k] = v 
}



// 2. 读写锁
type MapTplResCode struct {
  Data map[string]int
  Lock *sync.RWMutex
}


func (d MapTplResCode) Get(k string) (int, bool) {
  d.Lock.RLock()

  defer d.Lock.RUnlock()

  if v, ok := d.Data[k]; ok {
    return v, true
  }

  return 0, false
}


func (d MapTplResCode) Set(k string, v int) {
  d.Lock.Lock()

  defer d.Lock.Unlock()

  d.Data[k] = v
}


func (d MapTplResCode) Init() {
  d.Lock.Lock()

  defer d.Lock.Unlock()

  for key, _ := range d.Data {
    delete(d.Data, key)
  }
}































































