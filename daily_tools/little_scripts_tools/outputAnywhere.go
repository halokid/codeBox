/**
file:   output string to file anywhere use other language
build:  go build outputAnywhere.go
 */
package main

import (
	"os"
	"fmt"
)

var (
  destFile = "/tmp/outputAnywhere.txt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: string")
		os.Exit(0)
	}

	var s string
  s = string(os.Args[1]) + "\n"
  f, _:= os.OpenFile(destFile , os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
  f.WriteString(s)
  f.Close()
}



