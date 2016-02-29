package main 

import (
  "database/sql"
  "fmt"
  _ "github.com/ziutek/mymysql/godrv"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "net/http"
)

var MAX_POOL_SIZE = 20 
var MySQLPool chan *sql.DB


func getMySQL() *sql.DB {
  if MySQLPool == nil {
    MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
  }
  
  if len(MySQLPool) == 0 {
    go func() {
      for i := 0; i < MAX_POOL_SIZE/2; i++ {
        fmt.Println("crean DB conn....")
        mysqlc, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*test/root/")
        if err != nil {
          panic(err)
        }
        putMySQL(mysqlc)
      }
    } () 
  }
  return <-MySQLPool
}

func putMySQL(conn *sql.DB) {
  // fmt.Println("crean DB conn....")
  if MySQLPool == nil {
    MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
  }
  if len(MySQLPool) == MAX_POOL_SIZE {
    conn.Close()
    return
  }
  MySQLPool <-conn
}


func startHttpServer() {
  http.HandleFunc("/pool", pool)
  err := http.ListenAndServe(":9090", nil) 
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
    fmt.Println(123)
  }
}


func pool(w http.ResponseWriter, r *http.Request) {
  dbx := getMySQL()
  id := 1
  // var query = "select * from user_pwd"
  var query = "SELECT email from users WHERE id = ?"
  rows, err := dbx.Query(query, id)
  // rows, err := dbx.Query(query, "")
  if err != nil {
    log.Fatal(err)
  }
  
  var email string
  for rows.Next() {
    if err := rows.Scan(&email); err != nil {
      log.Fatal(err)
    }
    fmt.Println("Email address: ", email)
  }
  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }
  defer putMySQL(dbx)
}


func main() {
  startHttpServer()
}









