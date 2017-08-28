package utils

import (
  "fmt"
  "os"
)


func CheckErr(s string, err error) {
  if err != nil {
    fmt.Println(s)
    os.Exit(0)
  }
}