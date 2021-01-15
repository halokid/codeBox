package main

import (
  "log"
  "time"
)

func genMap() map[int]int {
  m := make(map[int]int)
  for i := 0; i < 1000000; i++ {
    m[i] = i
  }
  return m
}

func readMap(m map[int]int, i int) int {
  if v, ok := m[i]; ok {
    log.Println("存在map key", i)
    return v
  } else {
    log.Println("不存在map key", i)
  }
  return 0
}

func writeMap(m map[int]int) {
  for i := 0; i < 100; i++ {
    m[i] = i
  }
}

func main() {
  m := genMap()
  for i := 0; i < 100000; i++ {
    go readMap(m, i)      // todo: 如果是map数据没有写入，协程并发读取不会出现错误
    go writeMap(m)      // todo: 协程并发写入, 报错 concurrent map writes
  }

  time.Sleep(15 * time.Second)
}




