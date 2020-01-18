package json

import (
  "encoding/json"
  "fmt"
  "testing"
)

var jsonStr = `{
        "basic_info": {
          "name": "Mike",
          "age":  30
        },

        "job_info": {
          "skills": ["java", "GO", "PHP"]
        }
}`

func TestEmbededJson(t *testing.T) {
  e := new(Employee)
  // 把字符串先赋值结构体
  err := json.Unmarshal([]byte(jsonStr), e)
  if err != nil {
    t.Error(err)
  }

  fmt.Println(*e)
  fmt.Printf("%+v\n", *e)
  // 从结构体解码为 json
  if v, err := json.Marshal(e); err == nil {
    fmt.Println(string(v))
  } else {
    t.Error(err)
  }
}

func TestEasyJson(t *testing.T) {
  e := Employee{}
  // 把字符串先赋值结构体
  e.UnmarshalJSON([]byte(jsonStr))
  fmt.Println(e)
  // 从结构体解码为 json
  if v, err := e.MarshalJSON(); err!= nil {
    t.Error(err)
  } else {
    fmt.Println(string(v))
  }
}


func BenchmarkEmbeddedJson(b *testing.B) {
  b.ResetTimer()
  e := new(Employee)
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal([]byte(jsonStr), e)
    if err != nil {
      b.Error(err)
    }
    if _,err = json.Marshal(e); err != nil {
      b.Error(err)
    }
  }
}

func BenchmarkEasyJson(b *testing.B) {
  b.ResetTimer()
  e := Employee{}
  for i := 0; i < b.N; i++ {
    err := e.UnmarshalJSON([]byte(jsonStr))
    if err != nil {
      b.Error(err)
    }
    if _, err = e.MarshalJSON(); err != nil {
      b.Error(err)
    }
  }
}




