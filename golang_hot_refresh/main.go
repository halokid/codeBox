package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  r := gin.Default()
  r.GET("/hello", func(c *gin.Context) {
    c.String(http.StatusOK, "hello fresh reload")
  })

  _ = r.Run()
}
