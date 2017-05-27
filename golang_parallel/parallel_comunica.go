package main


func Generate(ch chan<- int) {
  for i := 2; ; i++ {
    ch <- i
  }
}


/**
in <-chan int 意思为把channel输入到in， 所以  in <- chan

out chan<- int, 意思为输出到这个 channel， 所以 out chan<-
 **/

func Filter(in <-chan int, out chan<- int, prime int) {
  for {
    println("-------- filter ------------")
    i := <-in
    println("i: ", i, "-------------")
    println("prime: ",  prime, "-------------")
    if i%prime != 0 {
      out <- i
    }
  }
}


func main() {
  ch := make(chan int)  // make the channel for ch
  go Generate(ch)   // generate ch for in， 这里改变一次ch的值为按顺序的自然数
  
  for i := 0; i < 10; i++ {
    prime := <-ch   // 事实上改变 ch 的值的代码段是在这行下面的，但是输出 prime 的时候，却已经是改变了的ch的值，所以 channel的原理就是这样，是独立于整个程序之外的
    println(prime, "\n")
    
    ch1 := make(chan int)
    go Filter(ch, ch1, prime)
    // ch = ch1
    
  }
}







