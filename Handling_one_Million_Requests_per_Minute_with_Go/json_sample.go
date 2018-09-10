package main 

import (
  "fmt"
  "encoding/json"
  "github.com/bitly/go-simplejson" // for json get
)


type  MyData struct {
  name        string `json"item"`
  other       string `json:"amount"`
}


/**

这里需要注意的就是后面单引号中的内容。

`json:"item"`
这个的作用，就是Name字段在从结构体实例编码到JSON数据格式的时候，使用item作为名字。算是一种重命名的方式吧。

**/


func main() {
  var detail MyData

  detail.name = "1"
  detail.other = "2"

  js_body, err := json.Marshal(detail)
  if err != nil {
    panic(err.Error()) 
  } 

  fmt.Println(js_body)

  js, err := simplejson.NewJson(js_body)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println(js)
}




