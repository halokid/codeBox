package main 
/**

这样的写法初学者经常会遇到的，很危险！ 与Java的foreach一样，都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝。 就像想修改切片元素的属性：

**/

import (
  "fmt"
) 

type student struct {
  Name  string
  Age   int 
}

func paseStudent() {
  m := make(map[string]*student)

  stus := []student {
    {Name:    "zhou",   Age:  24},
    {Name:    "lee",   Age:  23},
    {Name:    "wang",   Age:  22},
  }  


  /**
  for _, stu := range stus{
    m[stu.Name] = &stu 
    fmt.Println("stu.Name ----------------", stu.Name)
    fmt.Println("&stu ------------------- ", &stu)
  }

  fmt.Println("m -----------------------", m)
  // 输出 
  // map[zhou:0xc42000a060 lee:0xc42000a060 wang:0xc42000a060]
  // 三个内存位置都一样的， 所以这个是错误的写法
  **/

  // 正确的写法
  for i := 0; i < len(stus); i++ {
    m[stus[i].Name] = &stus[i]
  }

  for k, v := range m {
    fmt.Println(k, " => ", v.Name, " --- ", v.Age)
  } 

}

func main() {
  paseStudent() 

}










