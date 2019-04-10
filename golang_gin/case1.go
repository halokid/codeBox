package main
/**
用gin框架的时候，一些变量的上下文关系和进程启动之后，某些变量不变的具体问题
 */
import (
  "fmt"
  "github.com/gin-gonic/gin"
  "strconv"
  "time"
)

func getTime() string {
  timestamp := time.Now().Unix()
  //xx := int(timestamp)
  xx := strconv.FormatInt(timestamp, 10)
  return xx
}

func retBd(c *gin.Context)  {
  fmt.Println("run ping3...")
  getT2 := getTimexx()
  c.JSON(200, gin.H{
      "message": getT2,
    })
}

func main() {
  r := gin.Default()

  // 会变
  r.GET("/ping", func(c *gin.Context) {
    fmt.Println("run ping...")
    c.JSON(200, gin.H{
      "message": getTime(),
    })
  })

  // 不会变
  getT := getTime()
  r.GET("/ping1", func(c *gin.Context) {
    fmt.Println("run ping1...")
    c.JSON(200, gin.H{
      "message": getT,
    })
  })

  // 不会变
  getT2 := getTimexx()
  fmt.Println("run ping2...")
  r.GET("/ping2", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": getT2,
    })
  })

  // 会变
  r.GET("/ping3", retBd)

  r.Run() // listen and serve on 0.0.0.0:8080
}





