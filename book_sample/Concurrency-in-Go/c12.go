package main

import (
  "github.com/labstack/gommon/log"
  "net"
  "sync"
  "time"
)

func connectToService() interface{} {
  // 进行网络请求
  time.Sleep(1 * time.Second)
  // 返回一个空的 struct，不占任何的内存位置
  return struct {}{}
}

func warmServiceConnCache() *sync.Pool {
  // 网络高性能处理的时候， 预热加载网络处理
  p := &sync.Pool{
    New:      connectToService,
  }

  for i := 0; i < 10; i++ {
    p.Put(p.New())
  }
  return p
}

func startNetworkDaemon() *sync.WaitGroup {
  // 模拟进行网络请求
  var wg sync.WaitGroup
  wg.Add(1)

  go func() {
    // 优化方式， 预热加载
    connPool := warmServiceConnCache()
    server, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
      log.Fatalf("cannot listen:  %v", err)
    }
    defer server.Close()
    wg.Done()

    // fixme: 建完网络服务监听之后， 下面就是模拟有客户端链接这个服务的逻辑， 所以上面是可以wg.Done() 的了
    for {
      conn, err := server.Accept()
      if err != nil {
        log.Printf("cannot accept connection:  %v", err)
        continue
      }

      // 这里是模拟连接到服务器之后， 服务器做的逻辑动作
      //connectToService()


      svcConn := connPool.Get()     // 进行网络服务逻辑处理
      // 处理完了，放入池里面
      connPool.Put(svcConn)

      //fmt.Println(conn, "")
      conn.Close()
    }

  }()

  return &wg
}







