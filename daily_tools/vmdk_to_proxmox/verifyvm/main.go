package main

import (
  "os"

  "github.com/Unknwon/goconfig"

  "../utils"
)

func main() {
  vms, err := goconfig.LoadConfigFile("../vms_all.txt")
  utils.CheckErr("cannot read vms file", err)

  sections := vms.GetSectionList()

  f, _ := os.OpenFile("./verify_log.txt", os.O_CREATE|os.O_APPEND, 0666)
  for _, sec := range sections {
    secIp, err := vms.GetValue(sec, "ip")
    utils.CheckErr("read vm ip error", err)

    var logOut string
    if utils.CheckVmIp(secIp) {
      logOut = "--------------" + secIp + " OK ----------- \n\r"
    } else {
      logOut = "--------------" + secIp + " FAIL ----------- \n\r"
    }

    f.WriteString(logOut)
  }

  f.Close()

}



