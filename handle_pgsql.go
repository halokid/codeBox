package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"  
)

func main() {
  db, err := sql.Open("postgres", "user=postgres_user password=password host=172.16.2.29 dbname=my_postgres_db sslmode=disable")
  checkErr(err)
  
  rows, err := db.Query("select type, color, location from pg_equipment")
  checkErr(err)
  
  for rows.Next() {
    // var e_id int 
    var typeStr string
    var color string
    var location string
    // var dateTime string
    err = rows.Scan(&typeStr, &color, &location)
    checkErr(err)
    fmt.Println(typeStr)
    fmt.Println(color)
    fmt.Println(location)
    // fmt.Println(dateTime)
  }
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}