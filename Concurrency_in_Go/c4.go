package main

func main() {
  doneChan := make(chan string)

  go func() {
    // do some work
    doneChan <- "i am all done!"
  }()

  <- doneChan
}


