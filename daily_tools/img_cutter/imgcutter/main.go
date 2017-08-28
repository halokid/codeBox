package main

import (
  "runtime"
  "fmt"

  "github.com/disintegration/imaging"

  "../utils"
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  img, err := imaging.Open("./big.jpg")


  utils.CheckErr("img open error", err)

  dstimg := imaging.Resize(img, 200, 0, imaging.Box)

  err = imaging.Save(dstimg, "./resized.jpg")

  utils.CheckErr("resized err", err)

  fmt.Println("OK")
}



