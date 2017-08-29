/**
file:   output string to file anywhere use other language
build:  go build outputAnywhere.go
 */
package main

import "os"

var (
  destFile = "/tmp/outputAnywhere.txt"
)

func main() {
  get := []byte(os.Args[1])
  f, _:= os.OpenFile( destFile, os.O_APPEND|os.O_CREATE, 0777)
  f.Write(get)
  f.Close()
}



