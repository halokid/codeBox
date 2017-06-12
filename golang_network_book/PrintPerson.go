package  main

import (
  "html/template"
  "os"
  "fmt"
)

type Person struct {
  Name string
  Age int
  Emails []string
  Jobs []*Job     //指针的引用， 引用到 job 的结构体去
}


type Job struct {
  Employer string
  Role     string
}

const templ  = `The name is {{.Name}}.
        The age is {{.Age}}.
        {{range .Emails}}
            An email is {{.}}
        {{end}}

        {{with .Jobs}}
          {{range .}}
            An employer is {{.Employer}}
            and the role is {{.Role}}
          {{end}}
        {{end}}
`


func main() {
  job1 := Job{Employer:"jimmy", Role:"founder"}
  job2 := Job{Employer:"tom", Role:"gooder"}

  person := Person{
    Name:  "jan",
    Age:   50,
    Emails:  []string{"xxx@11.com", "sss@dd.com"},
    Jobs:   []*Job{&job1, &job2},
  }

  t := template.New("person template")
  t, err := t.Parse(templ)
  checkError(1, err)

  err = t.Execute(os.Stdout, person)
  checkError(2, err)
}

func checkError(code int, err error) {
  if err != nil {
    fmt.Println(err)
    os.Exit(code)
  }
}


























