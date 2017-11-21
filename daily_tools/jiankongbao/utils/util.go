package utils

import (
  "os"
  "fmt"
  "crypto/md5"
)


//md5加密
func SetMd5(s string) string {
  data := []byte(s)
  mds := md5.Sum(data)
  md5str := fmt.Sprintf("%x", mds)
  return md5str
}

//检查错误的函数
func CheckErr(s string, err error) {
  if err != nil {
    fmt.Println(err)
    fmt.Println(s)
    os.Exit(0)
  }
}


//匹配ums的告警类型
//监控宝消息状态：1为故障，2为提醒。  1匹配ums的紧急， 2匹配ums的警告
func SwiWarnStatus(code int) int {
  switch code {
    case 1:
      return 99
    case 2:
      return 66
    default:
      return 66
  }
}




