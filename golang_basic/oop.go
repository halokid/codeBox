package main 

import (
  "fmt"
)

//declare a object
type Rect struct {
  x, y float64
  width, height float64
}

//realize the func of the object
func (r *Rect) Area() float64 {
  return r.width * r.height
}


// golang has no construct func
// 对象的创建通常交由一个全局的创建函数 NewXXX 来命名 表示构造函数
func newRect(x, y, width, height float64) *Rect {
  return &Rect {x, y, width, height}
}

// go也提供了继续 但是财通了组合的玩法 称之为匿名组合
// 匿名组合 示例

type Base struct {
  Name string
}

func (base *Base) Foo() {
  base.Name = "Base Foo2"
}

func (base *Base) Bar() {
  base.Name = "Base Bar"
}

//==================================


type Foo struct {
  Base    // 这里声明这个结构体继续自 Base 结构体， 所以这个结构体拥有Base的一切方法和属性
  Name1 string
}

func (foo *Foo) Bar() {
  foo.Base.Bar()
}


func main() {
  foo := &Foo{}
  foo.Foo()
  fmt.Println(foo.Name)
  
  rect:=new(Rect)
  rect1:=&Rect{width:109,height:10}
  rect2:=&Rect{1,2,3,4}
  rect.width=19.9
  rect.height=22.1
  fmt.Println(rect.Area())
  fmt.Println(rect1.Area())
  fmt.Println(rect2.Area())
}















