package main

import "fmt"

type Programer interface {
  WriteHelloWorld()  string
}


type GoProgramer struct {

}

func (g *GoProgramer) WriteHelloWorld() string {
  return "何谓鸭子类型？就是这个函数名 WriteHelloWorld 一开始我不知道是什么函数，" +
         "但是名字跟上面的 Programer interface 定义的方法名是一样的, 那么我们就认为这个函数就是上面定义的接口" +
         "就相当于我们知道鸟是啥样的， 但是第一次看见鸭子，不知道鸭子是什么鬼，但是看上去又比较像鸟，所以我们也认为鸭子就是鸟"
}

func main() {
  var p Programer
  //fmt.Println(p.WriteHelloWorld())        # 定义错误，类型声明只是一个定义，并没有内存空间，所以不能应用到方法

  p = new(GoProgramer)        // GoProgramer是定义了一个内存空间， 而且也承接了 Program 的接口
  fmt.Println(p.WriteHelloWorld())

  px := new(GoProgramer)
  fmt.Println(px.WriteHelloWorld())
}
