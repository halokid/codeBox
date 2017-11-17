package main

import (
  "fmt"
  "net/http"

  "../utils"
)

//监控宝回调url处理
func jkbCallback(w http.ResponseWriter, r *http.Request)  {
  r.ParseForm()
  post_daa := r.Form
  fmt.Println(post_daa)

  msg_id := post_daa["msg_id"]
  //fmt.Println(msg_id)
  task_id := post_daa["task_id"]
  fault_time := post_daa["fault_time"]
  token := post_daa["token"]
  mytoken := "xxxxxxxxxxxxxxxxxxxxxx"

  check_token := utils.SetMd5(string(msg_id[0]) + string(task_id[0]) + string(fault_time[0]) + string(mytoken))
  if (check_token != token[0]) {
    panic("check token fail")
  } else {
    handleCallback()
  }
}


//处理callback
func handleCallback()  {

}


//主函数
func main() {
  http.HandleFunc("/", jkbCallback)
  http.ListenAndServe(":8088", nil)
}





