package reflect

import (
  "fmt"
  "reflect"
  "testing"
)

func TestTypeAndValue(t *testing.T) {
  var f int64 = 10
  t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
  t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
  t := reflect.TypeOf(v)
  switch t.Kind() {
  case reflect.Float32, reflect.Float64:
    fmt.Println("Float")
  case reflect.Int, reflect.Int32, reflect.Int64:
    fmt.Println("Integer")
  default:
    fmt.Println("Unknow", t)
  }
}

func TestBasicType(t *testing.T) {
  var f float64 = 12
  CheckType(f)
  CheckType(&f)
}

type Employee struct {
  EmployeeID        string
  Name              string  `format:"normal"`
  Age               int
}

func (e *Employee) UpdateAge(newVal int) {
  e.Age = newVal
}

type Customer struct {
  CookieID        string
  Name            string
  Age             int
}

/**
TypeOf 的 FieldByName  是返回属性的结构的
ValueOf 的 FieldByName  是返回值本身的
 */

func TestInvokeByName(t *testing.T) {
  e := &Employee{"1", "Mike", 30}
  t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
  t.Log("获取Name的值:", reflect.ValueOf(*e).FieldByName("Name"))
  if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
    t.Error("获取 `Name` 属性失败")
  } else {
    t.Log("属性Name 的结构为：", nameField, " Tag:format 的值为: ", nameField.Tag.Get("format"))
  }

  reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
  t.Log("Update Age: ", e)
}























