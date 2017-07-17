package utils

import (
  "os"
  "os/exec"
  "fmt"
  "strings"
  "math/rand"
  "strconv"
  //"sync"
)

/**
get new create vm id
 */
func GetNewvmId() string{
  vmids := exec.Command("ls")
  ids, _ := vmids.Output()
  //fmt.Println(string(ids))
  idsSli := strings.Split(string(ids), "\n")
  //fmt.Println(idsSli)

  newId := rand.Intn(10000)
  newIdStr := strconv.Itoa(newId)
  newIdStrMake := CheckVmId(newIdStr, idsSli)   //int type
  //fmt.Println(newIdStrMake)

  return newIdStrMake
}


/**
recheck vm id number
 */
func CheckVmId(newid string, existsIds []string) string{
  for _, id := range existsIds {
    if newid == id {
      newIdRandom := rand.Intn(10000)
      newIdStr := strconv.Itoa(newIdRandom)

      CheckVmId(newIdStr, existsIds)
    }
  }
  //i, _ := strconv.Atoi(newid)
  //return i
  return newid
}


/**
run create vm command
 */
//func makeVm(comm string, lock *sync.Mutex) {
//func makeVm(newVmId, comm string, wg *sync.WaitGroup) {
func MakeVm(newVmId, comm string) {
//func makeVm(comm string) {
  //lock := &sync.Mutex{}
  //lock.Lock()
  //defer wg.Done()

  //cmd := exec.Command(comm)
  cmd := exec.Command("/bin/bash", "-c", comm)
  out, _ := cmd.CombinedOutput()
  fmt.Println(string(out))
  f, _ := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND, 0666)
  logOut := "------------------ creating " + newVmId + " --------------------\n\r" + string(out) + "\n\r"
  fmt.Println(logOut)
  f.WriteString(logOut)
  f.Close()

  //lock.Unlock()
}

/**
check vm ip is reachable
 */
func CheckVmIp(ip string) bool {
  cmd := exec.Command("/bin/bash", "-c", "ping " + ip + " -c 1 -W 2")
  out, _ := cmd.CombinedOutput()
  if strings.Contains(string(out), "0% packet loss") {
    return true
  } else {
    return false
  }
  
}



func CheckErr(s string, err error) {
  if err != nil {
    fmt.Println(s)
    os.Exit(0)
  }
}



