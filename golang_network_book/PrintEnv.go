package main

import (
  "net/http"
  "fmt"
  "os"
)

func main() {
  fileServer := http.FileServer(http.Dir("./www"))
  http.Handle("/", fileServer)

  http.HandleFunc("/cgi/printenv", printEnv)

  err := http.ListenAndServe(":8000", nil)
  checkError(1, err)
}

func checkError(code int, err error) {
  if err != nil {
    fmt.Println("error")
  }
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
  env := os.Environ()
  writer.Write([]byte("<h1>env</h1>\n<pre>"))
  for _, v := range env {
    writer.Write([]byte(v + "\n"))
  }
  writer.Write([]byte("</pre>"))
}




