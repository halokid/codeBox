package main

import (
  _ "github.com/go-sql-driver/mysql"
  //"github.com/go-redis/redis"
  //"gopkg.in/redis.v4"
  "github.com/garyburd/redigo/redis"

  "database/sql"
  "fmt"
  "os"
  "time"
  "net/http"
)


type RunDaemon struct {
}

const (
  MYSQL_HOST = "172.20.71.173"
  MYSQL_USER = "root"
  MYSQL_PWD = "admin888"
  MYSQL_PORT = "3306"

  REDIS_HOST = "172.20.71.175"
  REDIS_PORT = "6381"

  SHOP = "http://172.20.71.174"
  DIS = "http://172.20.71.174"
  CACHECL = "http://172.20.71.174:8585"
)


func (runDaemon *RunDaemon) runMysql() {
  db, _ := sql.Open("mysql", MYSQL_USER + ":" + MYSQL_PWD +"@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT +")/test?charset=utf8")
  //db, _ := sql.Open("mysql", "root:admin888@tcp(172.20.71.173:3306)/test?charset=utf8")
  rows, err := db.Query("select username, sex from users")
  defer rows.Close()
  checkError("mysql connect", err)

 for rows.Next() {
   var username string
   var sex string
    err = rows.Scan(&username, &sex)
   checkError("rows err", err)
   fmt.Printf("username: %s, sex: %s\n", username, sex)
 }
}


func (runDaemon *RunDaemon) runRedis() {
  /**
  client := redis.NewClient(&redis.Options{
    Addr: "172.20.71.175:6379",
    Password: "",
    DB: 0,
  })

  pong, err := client.Ping().Result()
  fmt.Println(pong, err)


  errSet := client.Set("key1", "value1", 0).Err()
  checkError("redis set key", errSet)
  **/

  c, err := redis.Dial("tcp", REDIS_HOST + ":" + REDIS_PORT)
  checkError("redis conn", err)

  _, err = c.Do("SET", "pwdxxx4", "123456", "EX", "10")
  _, err = c.Do("SET", "pwdx", "123456", "EX", "10")
  _, err = c.Do("SET", "pwdasdaxx4", "123456", "EX", "10")
  _, err = c.Do("SET", "pwdxqwe123xx4", "123456", "EX", "10")
  _, err = c.Do("SET", "pwdxx123123114", "123456", "EX", "10")
  checkError("redis set key", err)
}


//visit web sometimes
func (runDaemon *RunDaemon) runWeb() {
  client := &http.Client{}
  request, err := http.NewRequest("GET", SHOP, nil)
  checkError("webclient error", err)
  response, err := client.Do(request)
  fmt.Println(response.Status)
  checkError("webclient do request", err)
}

func (runDaemon *RunDaemon) runWeb1() {
  client := &http.Client{}
  request, err := http.NewRequest("GET", DIS, nil)
  checkError("webclient error", err)
  response, err := client.Do(request)
  fmt.Println(response.Status)
  checkError("webclient do request", err)
}

func (runDaemon *RunDaemon) runWeb2() {
  client := &http.Client{}
  request, err := http.NewRequest("GET", CACHECL, nil)
  checkError("webclient error", err)
  response, err := client.Do(request)
  fmt.Println(response.Status)
  checkError("webclient do request", err)
}


func checkError(p string, err error) {
  if err != nil {
    fmt.Println(p)
    fmt.Println(os.Stderr)
    os.Exit(1)
  }
}

func main() {
  var runDaemon RunDaemon
  for {
    time.Sleep(10 * time.Second)
    runDaemon.runMysql()
    runDaemon.runRedis()
    runDaemon.runWeb()
    runDaemon.runWeb1()
    runDaemon.runWeb2()
  }
}









