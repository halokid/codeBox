package main

import (
  "reflect"
  "testing"
)

func add(a, b int) int {
  return a+b
}

func TestReflectFuncCall(t *testing.T) {
  // 将函数包装为反射值对象
  funcVal := reflect.ValueOf(add)

  // 构造函数参数， 传入两个整型值
  paramsList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

  // 反射调用函数
  retList := funcVal.Call(paramsList)

  // 获取第一个返回值， 取整型值
  t.Log(retList[0].Int())
}




