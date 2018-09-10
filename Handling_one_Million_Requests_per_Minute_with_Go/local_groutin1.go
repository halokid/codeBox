
/**
第一种方式是比较粗暴的， 直接接收到 http 请求
**/
func payloadHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return 
  }

  var content = &PayloadCollection()
  err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)) . Decode(&content)
  if err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
    return
  }


  for _, payload := range content.Payloads {
    go payload.UploadToS3()
  }

  w.WriteHeader(http.StatusOK)
}































