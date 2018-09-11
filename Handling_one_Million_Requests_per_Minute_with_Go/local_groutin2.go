
/**

第二种方法是经典的处理大请求的方法， 把HTTP请求先放进队列里面去，通过控制队列的容量， 使之不会产生很多HTTP请求
同时处理的情况， 所以web处理程序不会卡死，

常见的一种做法是，举例来说， 我们先把要处理的HTTP任务放在redis队列里面去， 队列里面储存我们要处理HTTP请求的
一些数据， 然后我们后台有个进程去实际读取redis队列， 进行逻辑处理

但是这个例子是用golang 的 channel 来做队列，实际上跟上面的情况是一样的

**/


var Queue = chan Payload

func init() {
  Queue = make(chan Payload, MAX_QUEUE)
}


func payloadHandler(w http.ResponseWriter, r *http.Request) {

  // 将包体读入到一个字符串进行json解析
  var content = &PayloadCollection()
  err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)) . Decode(&content)
  if err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // 遍历每个有效载荷， 逐个排队之后上传S3
  for _, payload := range content.Payloads {
    // 把 payload 入队列
    Queue <- payload

  }  
}

































