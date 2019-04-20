package main

import (
  "hello/utils"

  "fmt"
  "github.com/mattbaird/jsonpatch"
) 


var simpleA = `{"a":100, "b":200, "c":"hello"}`
var simpleB = `{"a":100, "b":200, "c":"goodbye"}`

func main() {
    patch, e := jsonpatch.CreatePatch([]byte(simpleA), []byte(simpleB))
    if e != nil {
        fmt.Printf("Error creating JSON patch:%v", e)
        return
    }
    for _, operation := range patch {
        fmt.Printf("%s\n", operation.Json())
    }

    utils.PrintText("Hi")
}


