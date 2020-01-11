package flexible

import (
  "errors"
  "reflect"
  "testing"
)

type Employee struct {
  EmployeeID    string
  Name          string `format:"normal"`
  Age           int
}


type Customer struct {
  CookieID      string
  Name          string
  Age           int
}

func TestDeepEqual(t *testing.T) {
  a := map[int]string{1: "one", 2: "two", 3: "three"}
  b := map[int]string{1: "one", 2: "two", 3: "three"}
  t.Log(reflect.DeepEqual(a, b))
  //fmt.Println(a == b)

  s1 := []int{1, 2, 3}
  s2 := []int{1, 2, 3}
  s3 := []int{2, 3, 1}
  t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
  t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))
  //fmt.Println(s1 == s2)

  c1 := Customer{"1", "Mike", 40}
  c2 := Customer{"1", "Mike", 40}
  //fmt.Println(c1 == c2)
  t.Log(c1 == c2)
  t.Log(reflect.DeepEqual(c1, c2))
}

func fillBySetting(st interface{}, settings map[string]interface{}) error {
  if reflect.TypeOf(st).Kind() != reflect.Ptr {       // 如果类型不是指针
    if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {   // 如果类型不是结构
      return errors.New("第一个参数应该是一个指向结构体的指针")
    }
  }

  if settings == nil {
    return errors.New("settings 不能是nil")
  }

  var (
    field     reflect.StructField
    ok        bool
  )
  for k, v := range settings {
    if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
      continue
    }

    if field.Type == reflect.TypeOf(v) {
      // 获取 st 的反射引用， 也就是 传入的这个  interface{} 的本身
      vstr := reflect.ValueOf(st)
      // 因为 传入的 st 是一个指针， 所以要用 Elem() 方法屈获取 针对指向的结构体， 才可以用 FieldByName方法
      vstr = vstr.Elem()
      // FieldByName 从结构类型获得， 因为 st 是针对， 所以要用  vstr.Elem() 取得结构类型
      vstr.FieldByName(k).Set(reflect.ValueOf(v))
    }
  }
  return nil
}

func TestFillNameAndAge(t *testing.T) {
  settings := map[string]interface{}{"Name": "Mike", "Age": 40}
  e := Employee{}
  if err := fillBySetting(&e, settings); err != nil {
    t.Fatal(err)
  }
  t.Log(e)

  c := new(Customer)
  if err := fillBySetting(c, settings); err != nil {
    t.Fatal(err)
  }
  t.Log(c)
}




























