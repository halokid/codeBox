package main
/**
todo： 同一个协程，假如有channel被写入没读取，或者读取没写入，都会阻塞该协程
 */
import "fmt"

func testSimple() {
  intChan := make(chan int)

  go func() {
   intChan <-1        // 起另外一个协程来写入，主协程即使没有接收channel，主协程一样不会阻塞
  }()

  //intChan <-1     // 在主协程写入，如果没有在其他线程接收的话，就会deadlock

  value := <-intChan
  fmt.Println("value: ", value)
}

func main() {
  testSimple()
}
