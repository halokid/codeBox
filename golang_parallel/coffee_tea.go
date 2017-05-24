package main

import "time"

func IsReady(what string) {
  time.Sleep(3 * time.Second)
  println(what, "is ready")
}


/**
func main() {
  println("Let's go")
  IsReady("Coffee")
  IsReady("Tea")
  println("I'm done here")
}
**/


/**
func main() {
  println("Let's go")
  go IsReady("Coffee")
  go IsReady("Tea")
  println("I'm done here")
}
**/

func main() {
  println("Let's go")
  go IsReady("Coffee")
  go IsReady("Tea")
  println("I'm done here")
  
  time.Sleep(7 * time.Second)
}