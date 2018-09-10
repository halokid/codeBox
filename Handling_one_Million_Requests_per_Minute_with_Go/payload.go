
type PayloadCollection struct {
  WindowsVersion          string      `json:"version"`
  Token                   string      `json:"token"`
  Payloads                []Payload   `json:"data"`
}


type Payload struct {

}

/**
这个是 http 接收到请求之后， 实际进行处理的逻辑代码
**/
func (p *Payload) UploadToS3() error {
  storage_path := fmt.Sprintf("%v/%v", p.storageFolder, time.Now().UnixNano())

  bucket := S3Bucket

  b := new(bytes.Buffer)
  encodeErr := json.NewEncoder(b).Encode(payload)
  if encodeErr != nil {
    return encodeErr
  } 

  var acl = s3.Private
  var contentType = "application/octet-stream"

  return bucket.PutReader( storage_path, b, int64(b.Len()), contentType, acl, s3.Options{} )
}


