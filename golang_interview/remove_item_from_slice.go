package main 

import (
  "fmt"
)

/**

// wrong
func (s *Slice) Remove(value interface{}) error {
  for i, v := range *s {
    if isEqual(value, v) {
      if i == len(*s) - 1 {
        *s = (*s)[:i]
      }  else {
        *s = append((*s)[:i], (*s)[i + 2:]...)
      }

      return nil
    } 
  }

  return ERR_ELEM_NT_EXIST
}



// right
func (s *Slice) Remove(value interface{}) error {
  for i, v := range *s {
    if isEqual(value, v) {
      // 整合slice的前后， 去掉 i 索引的元素
      *s = append((*s)[:i], (*s)[i + 1:]...)
      return nil
    }
  }
  return ERR_ELEM_NT_EXIST
}

**/

func main() {
  i := 1
  // ++i
  i++

  fmt.Println(i)
}


