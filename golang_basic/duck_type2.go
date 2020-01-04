package main

import "fmt"

type Duck interface {
  // 鸭子原型
  WhoAmI() error
}


type Bird struct {
  // 开始不知道是鸟， 但是也实现了 鸭子的方法， 所以认为这个也是鸭子类的东西
}

func (b *Bird) WhoAmI() error {
  fmt.Println("其实我是一只鸟， 因为都是鸟的外形， 有点像鸭子， 所以可以由鸭子类型统一抽象出WhoAmI的方法")
  return nil
}


// ------------------------------------------


type Chicken struct {
  // 开始不知道是鸡， 但是也实现了 鸭子的方法， 所以认为这个也是鸭子类的东西

}

func (c *Chicken) WhoAmI() error {
  fmt.Println("其实我是一只鸡， 因为都是鸟的外形， 有点像鸭子， 所以可以由鸭子类型统一抽象出WhoAmI的方法")
  return nil
}



func main() {
  var d Duck
  d = &Bird{}       // 因为要具体执行鸟的方法， 所以把 鸭子 抽象为 鸟
  d.WhoAmI()

  d = &Chicken{}     // 因为要具体执行鸡的方法， 所以把 鸭子 抽象为 鸟
  d.WhoAmI()
}




