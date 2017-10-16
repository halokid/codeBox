
package main

func main() {
	runtime.GOMAXPROCS(1)
	listener, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn.(*net.TCPConn))
	}
}


func handle(server *net.TCPConn) {
	defer server.Close()
	
	client, err := net.Dial("tcp", "127.0.0.1:8849")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Close()

	go func() {
		defer server.Close()
		defer client.Close()

		buf := make([]byte, 2048)
		io.CopyBuffer(server, client, buf)
	}()

	buf := make([]byte, 2048)
	io.CopyBuffer(client, server, buf)

}



