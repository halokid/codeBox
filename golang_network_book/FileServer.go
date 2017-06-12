package main

import (
  "net/http"
  "fmt"
  "os"
)

func main() {
  fileServer := http.FileServer(http.Dir("./www"))

  err := http.ListenAndServe(":8000", fileServer)
  checkError(1, err)
}

func checkError(code int, err error) {
  if err != nil {
    fmt.Println("error")
    os.Exit(code)
  }
}
