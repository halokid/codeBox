package main

import (
  "testing"
  "fmt"

  "github.com/tealeg/xlsx"
  //"github.com/Unknwon/goconfig"

  //"../utils"
)

func Test_exls(t *testing.T) {
  execFileName := "./test.xlsx"
  xlFile, err := xlsx.OpenFile(execFileName)
  if err != nil {
    fmt.Println(".....")
  }

  for _, sheet := range xlFile.Sheet {
    for _, row := range sheet.Rows {
      /**
      for _, cell := range row.Cells {
        text := cell.String()
        fmt.Printf("%s\n", text)
      }
      **/
      fmt.Println(row.Cells[1])
    }
  }
}


func Test_writeIni(t *testing.T) {
  //vms, err := goconfig.LoadConfigFile("./vms_test.txt")
  //utils.CheckErr("cannot read vms file", err)

  //sections := vms.GetSectionList()
}






