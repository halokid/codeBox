package main

import (
  "strconv"
  "time"
)

func getTimexx() string  {
  timestamp := time.Now().Unix()
  //xx := int(timestamp)
  xx := strconv.FormatInt(timestamp, 10)
  return xx
}
