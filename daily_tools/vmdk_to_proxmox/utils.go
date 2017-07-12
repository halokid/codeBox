package main

import (
  "os/exec"
  //"fmt"
  "strings"
  "math/rand"
  "strconv"
)

/**
get new create vm id
 */
func getNewvmId() string{
  vmids := exec.Command("ls")
  ids, _ := vmids.Output()
  //fmt.Println(string(ids))
  idsSli := strings.Split(string(ids), "\n")
  //fmt.Println(idsSli)

  newId := rand.Intn(10000)
  newIdStr := strconv.Itoa(newId)
  newIdStrMake := checkVmId(newIdStr, idsSli)   //int type
  //fmt.Println(newIdStrMake)

  return newIdStrMake
}


/**
recheck vm id number
 */
func checkVmId(newid string, existsIds []string) string{
  for _, id := range existsIds {
    if newid == id {
      newIdRandom := rand.Intn(10000)
      newIdStr := strconv.Itoa(newIdRandom)

      checkVmId(newIdStr, existsIds)
    }
  }
  //i, _ := strconv.Atoi(newid)
  //return i
  return newid
}






