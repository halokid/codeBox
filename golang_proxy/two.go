
var pool = make(chan net.Conn, 100)

func borrow() (net.Conn, error) {
	select {
	case conn := <- pool:
		return conn, nil
	default:
		return net.Dial("tcp", "127.0.0.1:8849")
	}
}


func release(conn net.Conn) error {
	select {
	case pool <- conn:
		//return to pool
		return nil
	default:
		//pool is overflow
		return conn.Close()
	}
}

func handle(server *net.TCPConn) {
	defer server.Close()
	client, err := borrow()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer release(client)

	go func() {
		defer server.Close()
		defer release(client)

		buf := make([]byte, 2048)
		io.CopyBuffer(server, client, buf)
	}()

	buf := make([]byte, 2048)
	io.CopyBuffer(client, server, buf)

}