### golng mysql库连接池分析

#### 0x1 背景
golang的协程是好用，但是有时候瓶颈并不在语言上，而是在后面的数据源上面，例如我们常见的mysql，redis等，当一个后端服务很多请求的时候，语言是能hold得住，但是
mysql产生错误，比如  too many connection,  too many time_wait 等等这些，今天我们就分析一下怎么解决这种问题



#### 0x2 代码范例
请查看main.go， [链接](https://github.com/halokid/codeBox/tree/master/golang_db_redis_pool)
（有帮忙的话请start或者follow一下哦，谢谢）

#### 0x3 分析

> 只执行ini函数， 检查mysql的进程显示为（原有的mysql是没有进程在处理的）

没执行前

```

mysql> show processlist;
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
| Id | User            | Host             | db   | Command | Time | State                  | Info             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
|  4 | event_scheduler | localhost        | NULL | Daemon  | 2304 | Waiting on empty queue | NULL             |
|  9 | root            | 10.244.1.1:64000 | test | Sleep   | 1315 |                        | NULL             |
| 10 | root            | 10.244.1.1:64022 | test | Query   |    0 | starting               | show processlist |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
3 rows in set (0.01 sec)

```


执行后
```

mysql> show processlist;
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
| Id | User            | Host             | db   | Command | Time | State                  | Info             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
|  4 | event_scheduler | localhost        | NULL | Daemon  | 2284 | Waiting on empty queue | NULL             |
|  9 | root            | 10.244.1.1:64000 | test | Sleep   | 1295 |                        | NULL             |
| 10 | root            | 10.244.1.1:64022 | test | Query   |    0 | starting               | show processlist |
| 13 | root            | 10.244.1.1:52134 | test | Sleep   |   20 |                        | NULL             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
4 rows in set (0.00 sec)


```

可见执行 db.Ping() 之后， process多了一个 Sleep 的连接，就是放了一个连接进 连接池



运行
```
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(5)
```
两句之后， 连接池并没有改变， 可见上面的逻辑是在 数据库处理逻辑真实执行的时候才生效的


执行协程查询
```
mysql> show processlist;
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
| Id | User            | Host             | db   | Command | Time | State                  | Info             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
|  4 | event_scheduler | localhost        | NULL | Daemon  | 4397 | Waiting on empty queue | NULL             |
|  9 | root            | 10.244.1.1:64000 | test | Sleep   | 3408 |                        | NULL             |
| 10 | root            | 10.244.1.1:64022 | test | Query   |    0 | starting               | show processlist |
| 19 | root            | 10.244.1.1:54823 | test | Sleep   |  952 |                        | NULL             |
| 20 | root            | 10.244.1.1:54824 | test | Sleep   | 1104 |                        | NULL             |
| 47 | root            | 10.244.1.1:57906 | test | Sleep   |    0 |                        | NULL             |
| 48 | root            | 10.244.1.1:57909 | test | Sleep   |    0 |                        | NULL             |
| 49 | root            | 10.244.1.1:57912 | test | Sleep   |    0 |                        | NULL             |
| 50 | root            | 10.244.1.1:57907 | test | Sleep   |    0 |                        | NULL             |
| 51 | root            | 10.244.1.1:57908 | test | Sleep   |    0 |                        | NULL             |
| 52 | root            | 10.244.1.1:57913 | test | Sleep   |    0 |                        | NULL             |
| 53 | root            | 10.244.1.1:57911 | test | Sleep   |    0 |                        | NULL             |
| 54 | root            | 10.244.1.1:57910 | test | Sleep   |    0 |                        | NULL             |
| 55 | root            | 10.244.1.1:57915 | test | Sleep   |    0 |                        | NULL             |
| 56 | root            | 10.244.1.1:57914 | test | Sleep   |    0 |                        | NULL             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
15 rows in set (0.00 sec)

```

执行完在等待
```
mysql> show processlist;
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
| Id | User            | Host             | db   | Command | Time | State                  | Info             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
|  4 | event_scheduler | localhost        | NULL | Daemon  | 3931 | Waiting on empty queue | NULL             |
|  9 | root            | 10.244.1.1:64000 | test | Sleep   | 2942 |                        | NULL             |
| 10 | root            | 10.244.1.1:64022 | test | Query   |    0 | starting               | show processlist |
| 19 | root            | 10.244.1.1:54823 | test | Sleep   |  486 |                        | NULL             |
| 20 | root            | 10.244.1.1:54824 | test | Sleep   |  638 |                        | NULL             |
| 32 | root            | 10.244.1.1:56588 | test | Sleep   |   22 |                        | NULL             |
| 33 | root            | 10.244.1.1:56591 | test | Sleep   |   22 |                        | NULL             |
| 34 | root            | 10.244.1.1:56589 | test | Sleep   |   22 |                        | NULL             |
| 35 | root            | 10.244.1.1:56590 | test | Sleep   |   22 |                        | NULL             |
| 36 | root            | 10.244.1.1:56592 | test | Sleep   |   22 |                        | NULL             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
10 rows in set (0.00 sec)
```


协程执行完之后
```
mysql> show processlist;
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
| Id | User            | Host             | db   | Command | Time | State                  | Info             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
|  4 | event_scheduler | localhost        | NULL | Daemon  | 3941 | Waiting on empty queue | NULL             |
|  9 | root            | 10.244.1.1:64000 | test | Sleep   | 2952 |                        | NULL             |
| 10 | root            | 10.244.1.1:64022 | test | Query   |    0 | starting               | show processlist |
| 19 | root            | 10.244.1.1:54823 | test | Sleep   |  496 |                        | NULL             |
| 20 | root            | 10.244.1.1:54824 | test | Sleep   |  648 |                        | NULL             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
5 rows in set (0.00 sec)
```
我们发现最大连接控制在了10个， 执行完之后还有5个连接在保持着

**这里有一个很重要的问题，就是连接池的过期时间**

0x4 深入分析
我们把  db.SetConnMaxLifetime(15 * time.Second)， 连接池连接的生命周期设置为 15秒, 我们会发现15秒之后，连接池的连接都会断掉
```
mysql> show processlist;
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
| Id | User            | Host             | db   | Command | Time | State                  | Info             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
|  4 | event_scheduler | localhost        | NULL | Daemon  | 4987 | Waiting on empty queue | NULL             |
|  9 | root            | 10.244.1.1:64000 | test | Sleep   | 3998 |                        | NULL             |
| 10 | root            | 10.244.1.1:64022 | test | Query   |    0 | starting               | show processlist |
| 19 | root            | 10.244.1.1:54823 | test | Sleep   | 1542 |                        | NULL             |
| 20 | root            | 10.244.1.1:54824 | test | Sleep   | 1694 |                        | NULL             |
+----+-----------------+------------------+------+---------+------+------------------------+------------------+
5 rows in set (0.00 sec)

```

30秒之后再次查询数据库
```golang
time.Sleep(30 * time.Second)
  rows, err := db.Query("select name from users")
  fmt.Println("err -----", err)
  defer rows.Close()
  for rows.Next(){
    var name string
    rows.Scan(&name)
    fmt.Println("name---", name)
  }
```
这个时候发现程序会重新发起新的db连接



#### 总结：

**mysql服务端的连接生命周期**

还有一种请况就是，我们的程序的连接池生命周期设置大于mysql服务器的生命周期设置， 这个时候就会有一种请况，假如我们重复用连接池的连接，会产生
连接错误的问题，解决方法有两种：

1. 可以在程序里面设置生命周期时间小于mysql服务端的连接生命周期时间就可以了
2. 增加程序的重连(keepalive）机制，就是定时发送一个连接包服务端
关于第2点我们我们以后可以再发散来说，一般如果允许的话，用第一种方式即可。

```
mysql> show variables like 'mysqlx_wait_timeout';
+---------------------+-------+
| Variable_name       | Value |
+---------------------+-------+
| mysqlx_wait_timeout | 28800 |
+---------------------+-------+
1 row in set (0.00 sec)
```










