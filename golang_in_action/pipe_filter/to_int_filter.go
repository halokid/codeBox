package pipe_filter

import (
  "errors"
  "strconv"
)

var ToIntFiltereError = errors.New("输入的数据应该是字符串格式")

type ToINtFilter struct {

}

func NewToIntFilter() *ToINtFilter {
  return &ToINtFilter{}
}

func (tif *ToINtFilter) Process(data Request) (Response, error) {
  parts, ok := data.([]string)
  if !ok {
    return nil, ToIntFiltereError
  }

  ret := []int{}
  for _, part := range parts {
    s, err := strconv.Atoi(part)
    if err != nil {
      return nil, err
    }
    ret = append(ret, s)
  }
  return ret, nil
}



