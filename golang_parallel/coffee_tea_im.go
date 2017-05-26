package main


import "time"

func IsReady(what string, ch chan<- bool) {
	time.Sleep( time.Second * 6)
	println(what, "is ready!")
	ch<- true
}


func main() {
	ch := make(chan bool)
	println("let's go!")
	go IsReady("Coffee", ch)
	go IsReady("Tea", ch)

	for i := 0; i < 2; i++ {
		<-ch
	}
}


