package main

import (
  "html/template"
  "os"
  "fmt"
  "strings"
)

type Person struct {
  Name string
  Emails []string
}

const templ = `the name is {{.Name}}.
    {{range .Emails}}
      an email is "{{. | emailExpand}}"
    {{end}}
`



func main() {
  person := Person {
    Name:         "jan",
    Emails:       []string{"jan@xxx.com", "xxxx@yyy.com"},
  }

  t := template.New("person template")

  t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})

  t, err := t.Parse(templ)
  checkError(1, err)

  err = t.Execute(os.Stdout, person)
  checkError(1, err)
}


func checkError(code int, err error) {
  if err != nil {
    fmt.Println("error")
  }
}


func EmailExpander(args ...interface{}) string {
  ok := false
  var s string

  if len(args) == 1 {
    s, ok = args[0].(string)
  }

  if !ok {
    s = fmt.Sprint(args...)
  }

  substrs := strings.Split(s, "@")
  if len(substrs) != 2 {
    return s
  }

  return (substrs[0] + "at" + substrs[1])

}























