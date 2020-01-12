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
  err := json.Unmarshal([]byte(jsonStr), e)
  if err != nil {
    t.Error(err)
  }

  fmt.Println(*e)
  fmt.Printf("%+v\n", *e)
  if v, err := json.Marshal(e); err == nil {
    fmt.Println(string(v))
  } else {
    t.Error(err)
  }
}



