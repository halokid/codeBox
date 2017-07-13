package main

import (
  "fmt"
  //"os"

  "github.com/Unknwon/goconfig"
  //"time"
  //"sync"
  //"runtime"
  "sync"
)

func main() {
  // read vms from config file
  vms, err := goconfig.LoadConfigFile("./vms_all.txt")
  CheckErr("cannot read vms file", err)

  sections := vms.GetSectionList()
  //vmCommSli := []string{}      //comm slice

  var wg sync.WaitGroup


  for _, sec := range sections {
    sockets, err := vms.GetValue(sec, "sockets")
    CheckErr("get sockets error", err)
    //fmt.Println(sockets)

    cores, err := vms.GetValue(sec, "cores")
    CheckErr("get cores error", err)
    //fmt.Println(cores)

    mem, err := vms.GetValue(sec, "mem")
    CheckErr("get mem error", err)
    //fmt.Println(mem)

    mac, err := vms.GetValue(sec, "mac")
    CheckErr("get mac error", err)
    //fmt.Println(mac)

    disk, err := vms.GetValue(sec, "disk")
    CheckErr("get disk error", err)
    //fmt.Println(disk)

    newVmId := getNewvmId()


    vmComm := "qm create " + newVmId + " -net0 vmxnet3,bridge=vmbr0 -name " + sec + " -memory " +
              mem + " -sockets " + sockets + " -cores " + cores + " -bootdisk ide0 -ide0 local:" + newVmId +
              "/vm-" + newVmId  + "-disk-1.vmdk,size=" + disk + "G; " +
              "sleep 3; " +
              "qm set " + newVmId + " -net0 vmxnet3='" + mac + "',bridge=vmbr0; " +
              "sleep 3; " +
              "mkdir -p /var/lib/vz/images/" + newVmId + "; " +
              "sleep 3; " +
              "cp ./" + sec + ".vmdk /var/lib/vz/images/" + newVmId + "/vm-" + newVmId + "-disk-1.vmdk; " +
              "sleep 3; " +
              "qm start " + newVmId

    fmt.Println(sockets + cores + mem + disk + mac + newVmId)
    //vmComm := "dir"   //跑完命令需要3秒

    //vmCommSli = append(vmCommSli, vmComm)
    //fmt.Println(vmComm)
    //fmt.Println("---------------------------------------------------------------------------")

    //lock := &sync.Mutex{}
    //lock.Lock()
    //go makeVm(vmComm, lock)

    wg.Add(1)

    go makeVm(newVmId, vmComm, &wg)

    //go makeVm(vmComm)
    //runtime.Gosched()

    //lock.Unlock()
    //time.Sleep(time.Second * 3)       //如果这里设置了 1 秒， 那么还没有跑完命令行，这个程序已经结束了，那么整个main进程结束，
                                      // makeVm也会结束， 命令行的结果写不进log文件
  }

  //fmt.Println(vmCommSli)
  wg.Wait()

  fmt.Println("----------------- all finished -----------------")
  //check the new vm status

}



