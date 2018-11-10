package main 

import (
  "fmt"
) 

type Slice []int


func NewSlice() Slice {
  return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
  *s = append(*s, elem)
  fmt.Println(elem)
  return s
}


func main() {
  s := NewSlice()

  // defer s.Add(1).Add(2)    // 这句的结果是输出 1 3 2, 因为defer 的逻辑是输出最后一步的， 因为  s.Add(1).Add(2)  的最后一步逻辑其实是   .Add(2)， 所以  .Add(1)这一步不会延迟， 会先执行

  defer s.Add(1))       // 这句的结果是输出 1 3

  s.Add(3)
}



func deferDemo() error {
  err := createResource1()
  if err != nil {
    return ERR_REATE_RESOURCE1_FAILD
  }

  defer func() {
    if err != nil {
      deftoryResource1()
    }
  }()

  err = createResource2()
  if err != nil {
    return ERR_REATE_RESOURCE2_FAILD
  }
}

























