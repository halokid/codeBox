
/**
幼稚的协程方式

第一种方式是比较粗暴的， 直接接收到 http 请求, 然后调用处理逻辑函数处理

当收到每分钟 1百万 的POST 请求， 这段代码很快就奔溃了
**/
func payloadHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return 
  }

  // 将包体读入到一个字符串进行json解析
  var content = &PayloadCollection()
  err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)) . Decode(&content)
  if err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
    return
  }


  // 遍历每个负载， 将其条目逐一队列化， 以期把它们发到S3
  for _, payload := range content.Payloads {
    // FIXME:  TODO: 
    
    go payload.UploadToS3()       // <----- 千万别这么干
  }

  w.WriteHeader(http.StatusOK)
}































