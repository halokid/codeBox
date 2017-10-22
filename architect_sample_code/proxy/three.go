
var pool = make(chan net.Conn, 100)

type client struct {
	conn net.Conn
	inUse *sync.WaitGroup
}


func borrow() (clt *client, err error) {
	var conn net.Conn
	select {
	case conn = <- pool:
	default:
		conn, err = net.Dial("tcp", "127.0.0.1:8849")
	}

	if err != nil {
		return nil, err
	}

	clt = &client {
		conn: conn,
		inUse: &sync.WaitGroup{},	
	}
	return
}


func release(clt *client) error {
	clt.conn.SetDeadline(time.Now().Add(-time.Second))
	clt.inUse.Done()
	clt.inUse.Wait()

	select {
	case pool <- clt.conn:
		//return to pool
		return nil
	default:
		//pool is overflow
		return clt.conn.Close()
	}
}


func handle(server *net.TCPConn) {
	defer server.Close()
	clt, err := borrow()

	if err != nil {
		fmt.Println(err)
		return
	}

	clt.inUse.Add(1)
	defer release(clt)

	go func() {
		clt.inUse.Add(1)
		defer server.Close()
		defer release()

		buf := make([]byte, 2048)
		io.CopyBuffer(server, clt.conn, buf)
	}()
 
	buf := make([]byte, 2048)
	io.CopyBuffer(clt.conn, server, buf)

}












