package main

import (
	"bufio"
	"fmt"
	"net"
)

//用来记录所有的客户端连接
var ConnMap map[string]*net.TCPConn

func main() {
	var tcpAddr *net.TCPAddr
	ConnMap = make(map[string]*net.TCPAddr)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	tcpListener, _ :=  net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()

		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		//新连接加入map
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn			//以IP地址作为key
		go tcpPipe(tcpConn)					//加入协程来处理连接
	}
}


func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader()
	for {
		message, err :=  reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
		//这里返回消息改为了广播
		boradcastMessage(conn.RemoteAddr().String() + ":" + string(message))
	}
}



func boradcastMessage(message string) {
	b := []byte(message)

	//遍历所有客户端并发送消息
	for _, conn := range ConnMap {
		conn.Write(b)
	}
}















