package main

import (
  "fmt"
  "sync"
)

func setM(wg *sync.WaitGroup, lk *sync.RWMutex, setMemx *map[int]int, i int) {
  defer wg.Done()
  fmt.Println("i ---- ", i)
  (*setMemx)[i] = i * 2
}

func main() {
  setMemx := make(map[int]int)
  wg := sync.WaitGroup{}
  lk := sync.RWMutex{}
  
  for i := 0; i < 10; i++ {
    wg.Add(1)
    setM(&wg, &lk, &setMemx, i)
  }

  wg.Wait()
  fmt.Println("setMen len ----------", len(setMemx))
  fmt.Println("setMem done ... ")
}
