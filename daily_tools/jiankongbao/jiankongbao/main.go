package main

import (
  "fmt"
  "net/http"

  "../utils"
  "net/url"
  "strconv"
  //"os"
  "time"
)

//监控宝回调url处理
func jkbCallback(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  post_data := r.Form

  msg_id := post_data["msg_id"]
  //fmt.Println(msg_id)
  task_id := post_data["task_id"]
  fault_time := post_data["fault_time"]
  token := post_data["token"][0]
  mytoken := "649f4e478b67db45906e92e29bb51fe5"

  check_token := utils.SetMd5(string(msg_id[0]) + string(task_id[0]) + string(fault_time[0]) + string(mytoken))
  fmt.Println("\n\r---------------------------------------------------------------------------------")
  fmt.Println(post_data)
  if (check_token != token) {
    fmt.Println("传来的token应该是:  " + check_token)
    fmt.Println("token fail: " + token)
    //panic("check token fail")
  } else {
    handleCallback()
  }
}




//处理callback
var dbhost = "10.86.22.71"
var dbport = "3306"
var dbuser = "root"
var dbpassword = "mysql"
var dbname = "bsmdb"


func handleCallback() {

}

//func handleCallback(post_data map[string]map[int]string)  {
func handleCallbackOld(post_data url.Values) {
  /**
  有两个逻辑， 假如消息的状态是 告警， 那么就新添加进ums 的 当前告警数据表
             假如消息的状态是 已修复， 那么就把 ums 原来对应的当前告警数据表的数据修改到 历史告警数据表去
   */

  //post_data := v
  //如果消息类型是 告警
  //fmt.Println(post_data)
  //os.Exit(1)
  //map[msg_id:[10887878] fault_time:[1876556565] task_id:[266577] task_user_id:[2] task_name:[jimmy_test] server_id:[1234] token:[1d3564d0fd2a8c833160d4a940bfdba5] msg_status:[1] msg_type:[1] content:[不可用监测点( 不可用)] task_type:[load] task_summary:[118.144.76.75(2)] message_detail:[{"metric":"load_average_5","compare":"gt","value":"1","unit":""]]

  if (post_data["message_status"][0] == "1") {
    db := utils.DbConn(dbhost, dbport, dbuser, dbpassword, dbname)

    //INSERT INTO `bsmdb`.`alarm` (`alarmid`, `deviceid`, `AlarmLevel`, `AlarmTimes`, `CreateDate`, `resname`, `UpdateDate`, `AlarmType`, `IsConfirm`, `ConfirmDate`, `IsRestore`, `RestoreDate`, `LastSendTime`, `ObjType`, `ObjField`, `EventType`, `MgIp`, `part`, `ObjName`, `ObjLabel`, `ObjId`, `ObjDesc`, `AlarmField`, `CurrentStatus`, `AppendField`, `AppendInfo`, `CreateHistory`, `IsSendConfirmMail`, `IsSendRestoreMail`, `Xcount`, `FirstCommentaryTime`, `SendCount`, `ConfirmBy`, `IsSendNotice`, `DelFlag`, `LastModifyTime`, `sourceFlag`) VALUES ('180', '854', '99', '2', '1511145304561', 'Oracle@ps', '1511145399658', '0', '0', NULL, '0', NULL, '1511145304561', 'oracleractbspace', 'used_percent', '表空间异常', 'psdbsvc.infinitus.com.cn', NULL, 'D.5.1.1.2.48', '/资源视图/数据库/OracleRac/Oracle@ps(psdbsvc.infinitus.com.cn)/表空间/HJH_TEST', '904', NULL, '使用率: 100% > 90%', 'Used: 100%', '使用率: 100% > 90%', '', '2017-11-20 10:35:04,2017-11-20 10:36:39', '0', '0', '1', NULL, '1', NULL, '1', '0', '1511145399667', 'umsptsvr01/8.8.8.8');


    /**
    INSERT INTO `bsmdb`.`alarm` (`alarmid`, `deviceid`, `AlarmLevel`, `AlarmTimes`, `CreateDate`, `resname`, `UpdateDate`, `AlarmType`, `IsConfirm`, `ConfirmDate`, `IsRestore`, `RestoreDate`, `LastSendTime`, `ObjType`, `ObjField`, `EventType`, `MgIp`, `part`, `ObjName`, `ObjLabel`, `ObjId`, `ObjDesc`, `AlarmField`, `CurrentStatus`, `AppendField`, `AppendInfo`, `CreateHistory`, `IsSendConfirmMail`, `IsSendRestoreMail`, `Xcount`, `FirstCommentaryTime`, `SendCount`, `ConfirmBy`, `IsSendNotice`, `DelFlag`, `LastModifyTime`, `sourceFlag`) VALUES ('182', '802', '99', '24448', '1510127259676', 'jimmy_test', '1510910772714', '0', '0', NULL, '0', NULL, '1510127259676', '', 'Status', '', '172.21.51.36', '监控宝', '', '', '825', NULL, '不可用监测点( 不可用)', '状态:不可用,剩余比率:0.0 %', '状态: 不可用', '序列号:&nbsp;&nbsp;&nbsp;&nbsp;TM100099240468  <br>名称:&nbsp;&nbsp;&nbsp;&nbsp;TRAY5 (Bypass)  <br>厂商:&nbsp;&nbsp;&nbsp;&nbsp;FUJI XEROX  <br>型号:&nbsp;&nbsp;&nbsp;&nbsp;FUJI XEROX External Media Handler  <br>剩余比率:&nbsp;&nbsp;&nbsp;&nbsp;0.0  %<br>剩余:&nbsp;&nbsp;&nbsp;&nbsp;0  张<br>容量:&nbsp;&nbsp;&nbsp;&nbsp;100  张<br>描述:&nbsp;&nbsp;&nbsp;&nbsp;Tray 5  <br>', '2017-11-17 08:38:05,2017-11-17 08:43:25', '0', '0', '1', NULL, '1', NULL, '1', '0', NULL, '');
     */


    //level_int, err := strconv.ParseInt(post_data["message_type"][0], 10, 32)
    level_int, err := strconv.Atoi(post_data["message_type"][0])
    utils.CheckErr("strconv int error", err)
    alarm_level := utils.SwiWarnStatus(level_int)
    //fmt.Println(alarm_level)
    //sql := "insert into alarm set CreateDate=" + post_data["fault_time"][0] + ", resname='" + post_data["task_name"][0] + "', AppendField='" + post_data["content"][0] + "', AlarmLevel=" + strconv.Itoa(alarm_level) + ", part='监控宝', DelFlag=0;"
    sql := "INSERT INTO `bsmdb`.`alarm` (`alarmid`, `deviceid`, `AlarmLevel`, `AlarmTimes`, `CreateDate`, `resname`, `UpdateDate`, `AlarmType`, `IsConfirm`, `ConfirmDate`, `IsRestore`, `RestoreDate`, `LastSendTime`, `ObjType`, `ObjField`, `EventType`, `MgIp`, `part`, `ObjName`, `ObjLabel`, `ObjId`, `ObjDesc`, `AlarmField`, `CurrentStatus`, `AppendField`, `AppendInfo`, `CreateHistory`, `IsSendConfirmMail`, `IsSendRestoreMail`, `Xcount`, `FirstCommentaryTime`, `SendCount`, `ConfirmBy`, `IsSendNotice`, `DelFlag`, `LastModifyTime`, `sourceFlag`) VALUES ('', '802', '" + strconv.Itoa(alarm_level) + "', '24448', '" + post_data["fault_time"][0] + "', '" + post_data["task_name"][0] + "', '" + post_data["fault_time"][0]  + "', '0', '0', NULL, '0', NULL, '" + post_data["fault_time"][0] + "', '', 'Status', '', '8.8.8.8', '监控宝', '', '', '825', NULL, '不可用监测点( 不可用)', '状态:不可用', '状态: 不可用', '序列号:xxx', '" + time.Now().Format("2006-01-02 15:04:05") + "', '0', '0', '1', NULL, '1', NULL, '1', '0', NULL, '');"

    //sql := `delete from alarm where alarmid=?`
    fmt.Println(sql)
    stmt, err := db.Prepare(sql)
    utils.CheckErr("insert prepare error", err)

    _, err = stmt.Exec()
    utils.CheckErr("insert exec error", err)


  } else if (post_data["message_status"][0] == "2") {
    fixCurrentWarn()
  }

}


//更改ums当前告警到历史告警
func fixCurrentWarn() {

}


//主函数
func main() {
  http.HandleFunc("/", jkbCallback)
  http.ListenAndServe(":8088", nil)
}





