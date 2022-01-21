
/*
networkd can not reach or server not running
 */

package main

import (
  "log"
  "net"
)

func main() {
  log.Println("begain dial...")
  conn, err := net.Dial("tcp", ":9999")
  if err != nil {
    log.Println("dial error ---", err)
    return
  }
  defer conn.Close()
  log.Println("dial OK")
}
