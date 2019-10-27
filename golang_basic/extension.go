package main
/**
  golang 的面向对象继承 和 定义
 */

import (
  "fmt"
)

type Pet struct {

}

func (p *Pet) Speak() {
  fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
  p.Speak()
  fmt.Println(" ", host)
}

type Dog struct {
  Pet
}

func (d *Dog) Speak() {
  fmt.Println("wong")
}

func main() {
  dog := new(Dog)
  dog.SpeakTo("xx")       // 这里并不会调用重新定义的 针对Dog 的 Speak 方法
}
