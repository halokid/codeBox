package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io"
  "io/ioutil"
  
  "github.com/gorilla/mux"
)



func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "welcome")
}


func TodoIndex(w http.ResponseWriter, r *http.Request) {
  // todos := Todos{
    // Todo{ Name: "write pressentation" },
    // Todo{ Name: "Host meetup" },
  // }
  
  
  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  
  if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
  }
}


func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Fprintln(w, "Todo show: ", todoId)
}



func TodoCreate(w http.ResponseWriter, r *http.Request) {
  var todo Todo
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    panic(err)
  }
  
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  
  //整个 HTTP POST 就是传输一个json过来，这里是解释JSON
  if err := json.Unmarshal(body, &todo); err != nil {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(422)    //unprocessble entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }
  
  
  // 把接收到的json 追加到 todo 列表里面去
  t := RepoCreateTodo(todo)
  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(t); err != nil {
    panic(err)
  }
  
}
















