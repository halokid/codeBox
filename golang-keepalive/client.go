package main


import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	b := []byte("time\n")
	conn.Write(b)

	<-quitSemaphore			//新开一个协程
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {	  //循环发包给服务端,这个是保持长连接的原因
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true		//如果产生错误就关闭协程
			break
		}
		time.Sleep(time.Second)			//每隔一秒发一个包
		b := []byte(msg)
		conn.Write(b)								//write动作就是发包
	}
}
