package main 

import (
  "fmt"
  "sync"
) 

type threadSafeSet struct {
  sync.RWMutex
  s []interface{}
}

func(set *threadSafeSet) Iter() <- chan interface{} (
  // ch := make(chan interface{})
  ch := make(chan interface{}, len(set.s)) 
  
  go func() {
    set.RLock()

    for elem, value := range set.s {
      ch <- elem
      fmt.Println("Iter: ", elem, value)
    }

    close(ch)
    set.RUnlock()
  } ()
  return ch
)

func main() {
  th := threadSafeSet{
    s:    []interface{}{"1", "2"}
  } 

  v := <- th.Iter()
  fmt.Sprintf("%s%v", "ch", v)
}


