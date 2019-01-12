// 热更新配置文件的程序范例
package main 

import (
  "os"
  "log"
  "sync"
  "syscall"
  "os/signal"
  "io/ioutil"
  "encoding/json"
)


type Config struct {
  Mode          string
  CacheSize     int 
}


var (
  config            *Config
  configLock   =    new(sync.RWMutex)       // 加锁
)



// load config
func loadConfig(fail bool) {
  file, err := ioutil.ReadFile("config.json")
  if err != nil {
    log.Println("open config: ", err)
    if fail { os.Exit(1) }
  }

  // todo: new 是返回一个指针地址的， 就是新定义的这个结构体的指针地址
  temp := new(Config)
  if err = json.Unmarshal(file, temp); err != nil {
    log.Println("parse config: ", err)
    if fail { os.Exit(1) }
  }

  configLock.Lock()
  config = temp
  configLock.Unlock()
}


func GetConfig() *Config {
  configLock.RLock()
  defer configLock.RUnlock()
  return config 
}


func init() {
  loadConfig(true) 
  s := make(chan os.Signal, 1)
  signal.Notify(s, syscall.SIGUSR2) 

  go func() {
    for {
      <- s
      loadConfig(false)
      log.Println("Reloaded")
    }
  }()
}













