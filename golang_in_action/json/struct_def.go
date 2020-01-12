package json

/**
golang 内置是用  发射机制来实现json的解释的， 类似 nameField.Tag.Get("format"))  FieldTag
 */

type BasicInfo struct {
  Name      string  `json:"name"`
  Age       int     `json:"age"`
}

type JobInfo struct {
  Skills    []string  `json:"skills"`
}

type Employee struct {
  BasicInfo   BasicInfo   `json:"basic_info"`
  JobInfo     JobInfo     `json:"job_info"`
}
