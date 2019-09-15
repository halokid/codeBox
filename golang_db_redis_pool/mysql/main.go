package main

import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/go-sql-driver/mysql"
  "sync"
  "time"
)

var db *sql.DB
var err error

func init() {
  db, err = sql.Open("mysql", "root:123456@tcp(192.168.1.129:30006)/test?charset=utf8")
  if err != nil {
    log.Fatalf("db conn err ----- ", err)
  }
  db.Ping()     // 先放一个连接进连接池
  if err != nil {
    log.Fatalf("db ping err")
  }

  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(5)
  db.SetConnMaxLifetime(15 * time.Second)
}

func GetUser() {
  fmt.Println("start GetUser...")
   rows, err := db.Query("select name, age from users")
   defer rows.Close()
   if err != nil {
     log.Fatal("GetUser sql err")
   }
   time.Sleep(2 * time.Second)
   for rows.Next() {
     var name string
     var age int
     rows.Scan(&name, &age)
     fmt.Println("name:", name, "--- age", age)
   }
}

func main() {
  fmt.Println("mysql连接池测试开始...")
  var wg sync.WaitGroup

  wg.Add(10)
  for i := 0; i < 10; i++ {
    go func() {
      defer wg.Done()
      GetUser()
    }()
  }

  wg.Wait()

  time.Sleep(30 * time.Second)
  rows, err := db.Query("select name from users")
  fmt.Println("err -----", err)
  defer rows.Close()
  for rows.Next(){
    var name string
    rows.Scan(&name)
    fmt.Println("name---", name)
  }

  time.Sleep(100 * time.Second)
}





