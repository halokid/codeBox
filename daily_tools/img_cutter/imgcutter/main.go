package main

import (
  "runtime"
  "fmt"

  "github.com/disintegration/imaging"

  "../utils"
  "os"
  "strconv"
)

func main() {

  if len(os.Args) < 3 {
    fmt.Println("Usage: $1 $2 $3, check readme")
    os.Exit(1)
  }

  runtime.GOMAXPROCS(runtime.NumCPU())

  img, err := imaging.Open(os.Args[1])


  utils.CheckErr("img open error", err)

  width, err := strconv.Atoi(string(os.Args[3]))
  utils.CheckErr("get with error", err)

  dstimg := imaging.Resize(img, width, 0, imaging.Box)

  err = imaging.Save(dstimg, os.Args[2])

  utils.CheckErr("resized err", err)

  fmt.Println("OK")
}



