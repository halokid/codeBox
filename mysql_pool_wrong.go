
package main

import(
  
)

var MySQLPool chan *mysql.MySQL

func getMySQL() *mysql.MySQL {
  if MySqlPool == nil {
    MySqlPool = make(chan *mysql.MySQL, MAX_POOL_SIZE)
  }
  if len(MySqlPool) == 0 {
    go func() {
      for i := 0; i < MAX_POOL_SIZE/2; i++ {
        mysql := mysql.New()
        err :=  mysql.Connect("127.0.0.1", "root", "", "wgt", 3306)
        if err != nil {
          panic(err.String())
        }
        putMySQL(mysql)
      }
    } ()
  }
  return <-MySQLPool
}

func putMySQL(conn *mysql.MySQL) {
  if MySQLPool == nil {
    MySQLPool = make(chan *mysql.MySQL, MAX_POOL_SIZE)
  }
  if  len(MySQLPool) == MAX_POOL_SIZE {
    conn.Close()
    return
  }
  MySQLPool <-conn
}


