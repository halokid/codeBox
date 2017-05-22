package main

import (
  "fmt"
  // "math/rand"
)


type query struct {
  sql chan string
  result chan string
}


func execQuery(q query) {
  go func() {
    sql := <-q.sql
    
    q.result <- "get " + sql
  }()
}


func main() {
  q := query{make(chan string, 1), make(chan string, 1)}
  
  execQuery(q)
  
  q.sql <- "select * from table"
  
  fmt.Println(<-q.result)
}






































