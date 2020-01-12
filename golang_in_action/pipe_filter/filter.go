package pipe_filter


type Request interface {
  // 统一封装输入数据格式

}

type Response interface {
  // 统一封装输出数据格式
}


type Filter interface {
  Process(data Request) (Response, error)
}




