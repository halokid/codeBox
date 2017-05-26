package main


import (
	"regexp"
	"strconv"
)

func echo(in chan string, out chan string) {
	out <- (<-in + "\n")
}


func sed(in chan string, out chan string) {
	re := regexp.MustCompile("hello")
	out <- re.ReplaceAllString(<-in, "goodbye")
}


func wc(in chan string, out chan string) {
	out <- strconv.Itoa(len(<-in))
}



func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go echo(ch1, ch2)
	go sed(ch2, ch3)
	go wc(ch3, ch4)

	ch1 <- "hello world"
	println(<-ch4)
	println(ch4)



}








