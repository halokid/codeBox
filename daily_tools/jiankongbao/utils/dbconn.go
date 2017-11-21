package utils

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

//返回数据库连接
func DbConn(host, port, user, password, database string) *sql.DB {
  conn := user + ":" + password + "@tcp(" + host + ":" + port +")/" + database + "?charset=utf8"
  db, err := sql.Open("mysql", conn)
  CheckErr("mysql conn error", err)
  return db
}
