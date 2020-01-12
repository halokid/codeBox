package pipe_filter

import (
  "errors"
  "strings"
)

var SplitFilterWrongFormatError = errors.New("输入的数据应该是字符串格式")


type SplitFilter struct {
  // 以文件名来定义一个struct， 通常是一种组织代码的方式
  delimiter  string
}

func NewSpliterFilter(delimiter string) *SplitFilter {
  // 定义一个struct 之后， 再定义一个创建 struct 的方法， 是一种组织代码的方式
  return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
  str, ok := data.(string)
  if !ok {
    return nil, SplitFilterWrongFormatError
  }
  parts := strings.Split(str, sf.delimiter)
  return parts, nil
}






