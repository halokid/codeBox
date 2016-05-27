package main

import (
	"fmt"
)


func deferFunc() int  {
	index := 0

	fc := func() {
		fmt.Println(index, "匿名函数1")
		index++

		defer func() {
			fmt.Println(index, "匿名函数1-1")
			index++
		}()
	}

		defer func() {
			fmt.Println(index, "匿名函数2")
			index++
		}()

		defer fc()

	 return func() int {
			fmt.Println(index, "匿名函数3")
			index++
			return index
	 }()

}

func main() {
	deferFunc()
}










