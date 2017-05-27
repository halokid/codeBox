package main

func main() {
  ch1 := make(chan int)
  ch2 := make(chan int)
  
  go func() {
    for {
      <-ch1
    }
  }()

  for {
    select {
      case ch1 <- 0:
        print(0)
      case ch1 <- 1:
        print(1)
      case ch2 <- 2:
        print(2)
    }
  }
  
  
  
}






