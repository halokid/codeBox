package main

import (
  "fmt"
  //"os"

  "github.com/Unknwon/goconfig"
  //"time"
  //"sync"
  //"runtime"

  "./utils"
  "strings"
)

func main() {
  // read vms from config file
  vms, err := goconfig.LoadConfigFile("./xlsxtoini/vms_all.txt")
  utils.CheckErr("cannot read vms file", err)

  sections := vms.GetSectionList()
  //vmCommSli := []string{}      //comm slice

  //var wg sync.WaitGroup


  for _, sec := range sections {
    /**
    sockets, err := vms.GetValue(sec, "sockets")
    utils.CheckErr("get sockets error", err)
    //fmt.Println(sockets)

    cores, err := vms.GetValue(sec, "cores")
    utils.CheckErr("get cores error", err)
    //fmt.Println(cores)
    **/
    cpu, err := vms.GetValue(sec, "cpu")
    utils.CheckErr("get cpu error", err)
    sockets := "1"
    cores   := cpu

    memStr, err := vms.GetValue(sec, "mem")
    memSli := strings.Split(".", memStr)
    mem := memSli[0]
    utils.CheckErr("get mem error", err)
    //fmt.Println(mem)

    mac1, err := vms.GetValue(sec, "mac1")
    utils.CheckErr("get mac error", err)
    //fmt.Println(mac)

    disk1, err := vms.GetValue(sec, "disk1")
    utils.CheckErr("get disk error", err)
    //fmt.Println(disk)

    newVmId := utils.GetNewvmId()

    mac2, mac2err := vms.GetValue(sec, "mac2")
    //utils.CheckErr("get mac2 error", err)

    ///**
    var mac2Comm string
    if mac2err == nil {
      mac2Comm = "qm set " + newVmId + " -net1 vmxnet3='" + mac2 + "',bridge=vmbr0;"
    } else {
      mac2Comm = ""
    }
    //**/


    disk2, disk2err := vms.GetValue(sec, "disk2")
    var disk2Comm string
    if disk2err == nil {
      disk2Comm = "qm set " + newVmId + " ide1 local:" + newVmId  + "/vm-" + newVmId  + "-disk-2.vmdk,size=" + disk2 + "G; " +
                  "mv ./" + sec + "-2.vmdk /var/lib/vz/images/" + newVmId + "/vm-" + newVmId + "-disk-2.vmdk; "
    } else {
      disk2Comm = ""
    }

		///**
    vmComm := "qm create " + newVmId + " -net0 vmxnet3,bridge=vmbr0 -name " + sec + " -memory " +
              mem + " -sockets " + sockets + " -cores " + cores + " -bootdisk ide0 -ide0 local:" + newVmId +
              "/vm-" + newVmId  + "-disk-1.vmdk,size=" + disk1 + "G; " +
              "sleep 3; " +
              "qm set " + newVmId + " -net0 vmxnet3='" + mac1 + "',bridge=vmbr0; " +
              "sleep 3; " +
              mac2Comm +
              "sleep 3; " +
              "mkdir -p /var/lib/vz/images/" + newVmId + "; " +
              "sleep 3; " +
              "mv ./" + sec + ".vmdk /var/lib/vz/images/" + newVmId + "/vm-" + newVmId + "-disk-1.vmdk; " +
              "sleep 3;" +
              disk2Comm +
              "sleep 3; " +
              "qm start " + newVmId
		//**/


    fmt.Println(sockets + cores + mem + disk1 + mac1 + newVmId)
    //vmComm := "dir"   //跑完命令需要3秒

    //vmCommSli = append(vmCommSli, vmComm)
    fmt.Println(vmComm)
    //fmt.Println("---------------------------------------------------------------------------")

    //lock := &sync.Mutex{}
    //lock.Lock()
    //go makeVm(vmComm, lock)

    //wg.Add(1)

    //go makeVm(newVmId, vmComm, &wg)
    //makeVm(newVmId, vmComm, &wg)
		//vmComm := "ls; sleep 10"
    utils.MakeVm(newVmId, vmComm)

    //go makeVm(vmComm)
    //runtime.Gosched()

    //lock.Unlock()
    //time.Sleep(time.Second * 3)       //如果这里设置了 1 秒， 那么还没有跑完命令行，这个程序已经结束了，那么整个main进程结束，
                                      // makeVm也会结束， 命令行的结果写不进log文件
  }

  //fmt.Println(vmCommSli)
  //wg.Wait()

  //check the new vm status

  fmt.Println("----------------- all finished -----------------")

}



