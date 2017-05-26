package main

import (
	"regexp"
	"strconv"
)


func echo (in chan string) chan string {
	out := make(chan string)
	go func() {
		out <- (<-in + "\n")
	}()

	return out
}


func sed(in chan string) chan string {
	out := make(chan string)
	re := regexp.MustCompile("hello")
	go func() {
		out <- re.ReplaceAllString(<-in, "goodbye")
	}()

	return out
}


func wc(in chan string) chan string {
	out := make(chan string)
	go func() {
		out <- strconv.Itoa(len(<-in))
	}()

	return out
}


func main() {
	in := make(chan string)
	out := wc(sed(echo(in)))
	in <- "hello world"
	println(<-out)
}











