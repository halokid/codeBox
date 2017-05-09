package main

import (
  "net/http"
  "github.com/gorilla/mux"
)



/** 没有装饰器的写法 --------
func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  
  for _, route := range routes {
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(route.HandlerFunc)
  }
  return router
}
**/


/** 有装饰器的写法 ---------------- **/
func NewRouter() *mux.Router {

  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler                  //declare handler
    
    handler = route.HandlerFunc               //set handler
    handler = Logger(handler, route.Name)     //装饰 handler 
    
    
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)  // 真正指定 Handler 的逻辑代码地方
  }
  return router
}







