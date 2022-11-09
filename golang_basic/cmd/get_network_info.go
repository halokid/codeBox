package main
import (
  "fmt"
  "github.com/safchain/ethtool"
  "os/exec"
  "strings"
)
func main() {
  baseNicPath := "/sys/class/net/"
  cmd := exec.Command("ls", baseNicPath)
  buf, err := cmd.Output()
  if err != nil {
    //fmt.Println("Error:", err)
    return
  }
  output := string(buf)
  for _, device := range strings.Split(output, "\n") {
    if len(device) > 1 {
      fmt.Println(device)
      ethHandle, err := ethtool.NewEthtool()
      if err != nil {
        panic(any(err.Error()))
      }
      defer ethHandle.Close()
      stats, err := ethHandle.LinkState(device)
      if err != nil {
        panic(any(err.Error()))
      }
      fmt.Printf("LinkName: %s LinkState: %d\n", device, stats)
    }
  }
}

