package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	//控制台聊天功能
	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "quit" {			//如果输入退出指令
			break
		}

		b := []byte(msg + "\n")
		conn.Write(b)
	}
}


func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)

		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}
