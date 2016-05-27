package main

import (
	"fmt"
)


func main() {
fc := func( msg string) {
	fmt.Println("you say :", msg)
}

fmt.Printf("%T \n", fc)

fc("hello, my love")

func(msg string) {
	fmt.Println("say :", msg)
}("I love to code")

}



