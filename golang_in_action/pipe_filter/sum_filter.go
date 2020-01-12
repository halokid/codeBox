package pipe_filter

import "errors"

var SumFilterWrongFormatError = errors.New("输入的格式应该是字符串")

type SumFilter struct {

}

func NewSumFilter() *SumFilter {
  return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
  elems, ok := data.([]int)
  if !ok {
    return nil, SumFilterWrongFormatError
  }

  ret := 0
  for _, elem := range elems {
    ret += elem
  }
  return ret, nil
}




