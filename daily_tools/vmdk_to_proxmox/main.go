package main

import (
  "fmt"
  "os"

  "github.com/Unknwon/goconfig"
)

func main() {
  // read vms from config file
  vms, err := goconfig.LoadConfigFile("./vms_all.txt")
  CheckErr("cannot read vms file", err)

  sections := vms.GetSectionList()
  for _, sec := range sections {
    sockets, err := vms.GetValue(sec, "sockets")
    CheckErr("get sockets error", err)
    fmt.Println(sockets)

    cores, err := vms.GetValue(sec, "cores")
    CheckErr("get cores error", err)
    fmt.Println(cores)

    mem, err := vms.GetValue(sec, "mem")
    CheckErr("get mem error", err)
    fmt.Println(mem)

    mac, err := vms.GetValue(sec, "mac")
    CheckErr("get mac error", err)
    fmt.Println(mac)

    disk, err := vms.GetValue(sec, "disk")
    CheckErr("get disk error", err)
    fmt.Println(disk)

    newVmId := getNewvmId()

    vmComm := "qm create " + newVmId + "-net0 vmxnet3,bridge=vmbr0 -name auto-make" + newVmId + " -memory " +
              mem + " -sockets " + sockets + " -cores " + cores + " -bootdisk ide0 -ide0 local:" + newVmId +
              "/vm-" + newVmId  + "-disk-1.vmdk,size=" + disk + "G;" +
              "qm set " + newVmId + " -net0 vmxnet3='" + mac + "',bridge=vmbr0"

    fmt.Println(vmComm)
    fmt.Println("---------------------------------------------------------------------------")
  }




  //use vms config  to create virtual-machine


  //check the new vm status


}


func CheckErr(s string, err error) {
  if err != nil {
    fmt.Println(s)
    os.Exit(0)
  }
}
