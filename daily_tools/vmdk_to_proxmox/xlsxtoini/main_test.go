package main

import (
  "testing"
  "fmt"

  "github.com/tealeg/xlsx"
  "github.com/Unknwon/goconfig"

  "../utils"
)


/**
// ---------------- qm vm command ------------------------
2280008052:71:54:AB:CD:E18081
qm create 8081 -net0 vmxnet3,bridge=vmbr0 -name auto-make1 -memory 8000 -sockets 2 -cores 2 -bootdisk ide0 -ide0 local:8081/vm-8081-disk-1.vmdk,size=80G; sleep 3; qm set 8081 -net0 vmxnet3='52:71:54:AB:CD:E1',bridge=vmbr0; sleep 3; qm set 8081 -net1 vmxnet3='52:71:54:AB:CD:E2',bridge=vmbr0;sleep 3;sleep 3; mkdir -p /var/lib/vz/images/8081; sleep 3; mv ./auto-make1.vmdk /var/lib/vz/images/8081/vm-8081-disk-1.vmdk; sleep 3; qm start 8081

------------------ creating 8081 --------------------


2280008052:71:54:AB:CD:E37887
qm create 7887 -net0 vmxnet3,bridge=vmbr0 -name auto-make2 -memory 8000 -sockets 2 -cores 2 -bootdisk ide0 -ide0 local:7887/vm-7887-disk-1.vmdk,size=80G; sleep 3; qm set 7887 -net0 vmxnet3='52:71:54:AB:CD:E3',bridge=vmbr0; sleep 3; sleep 3;qm set 7887ide1 local:7887/vm-7887-disk-2.vmdk,size=10G; sleep 3; mkdir -p /var/lib/vz/images/7887; sleep 3; mv ./auto-make2.vmdk /var/lib/vz/images/7887/vm-7887-disk-1.vmdk; sleep 3; qm start 7887

------------------ creating 7887 --------------------
 */

/**
func Test_exls(t *testing.T) {
  execFileName := "./test.xlsx"
  xlFile, err := xlsx.OpenFile(execFileName)
  if err != nil {
    fmt.Println(".....")
  }

  for _, sheet := range xlFile.Sheet {
    for _, row := range sheet.Rows {
      for _, cell := range row.Cells {
        text := cell.String()
        fmt.Printf("%s\n", text)
      }
      fmt.Println(row.Cells)
    }
  }
}
**/


func Test_writeIni(t *testing.T) {
  /**
  vms, err := goconfig.LoadConfigFile("./vms_test.txt")
  utils.CheckErr("cannot read vms file", err)

  vms.SetValue("auto-make1", "sockets", "2")
  err = goconfig.SaveConfigFile(vms, "./vms_test.txt")

  fmt.Println("run here")
  **/
  XlsToIni()
}


func XlsToIni() {
  execFileName := "./test.xlsx"
  xlFile, err := xlsx.OpenFile(execFileName)
  if err != nil {
    utils.CheckErr("read xls file error", err)
  }

  //read ini file
  vms, err := goconfig.LoadConfigFile("./vms_all.txt")
  utils.CheckErr("cannot read vms file", err)

  for _, sheet := range xlFile.Sheet {
    for i, row := range sheet.Rows {
      if i != 0 {
        //fmt.Println(row.Cells)
        //[auto-make5 10.86.21.193 7.9379999999999997 8 00:50:56:9a:ff:86 DPortGroup_10.86.21       200   ]
        item := row.Cells
        //fmt.Println(item[0])
        //secName := string(item[0])
        secName := fmt.Sprintf("%s", item[0])
        //fmt.Println(secName)
        vms.SetValue(secName, "ip", fmt.Sprintf("%s", item[1]))
        vms.SetValue(secName, "mem", fmt.Sprintf("%s", item[2]))
        vms.SetValue(secName, "cpu", fmt.Sprintf("%s", item[3]))
        vms.SetValue(secName, "mac1", fmt.Sprintf("%s", item[4]))
        vms.SetValue(secName, "mac2", fmt.Sprintf("%s", item[6]))
        vms.SetValue(secName, "disk1", fmt.Sprintf("%s", item[12]))
        vms.SetValue(secName, "disk2", fmt.Sprintf("%s", item[13]))
        /**
        secName := item[0]
        vms.SetValue(secName, "ip", item[1])
        vms.SetValue(secName, "mem", item[2])
        vms.SetValue(secName, "cpu", item[3])
        vms.SetValue(secName, "mac1", item[4])
        vms.SetValue(secName, "mac2", item[6])
        vms.SetValue(secName, "disk1", item[12])
        vms.SetValue(secName, "disk2", item[13])
        **/
      }
    }
  }
  err = goconfig.SaveConfigFile(vms, "./vms_all.txt")
  utils.CheckErr("write vms_all.txt error", err)
}






