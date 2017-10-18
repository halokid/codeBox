
var pool = make(chan net.Conn, 100)

type client struct {
	conn net.Conn							//net conn
	inUse *sync.WaitGroup			//sync lock
	isValid int32							//check valid or not
}


const maybeValid = 0
const isValid = 1
const isInvalid = 2

func (clt *client) Read(b []byte) (n int, err error) {
	n, err = clt.conn.Read(b)
	if err != nil {
		if !isTimeoutError(err) {
			atomic.StoreInt32(&clt.isValid, isInvalid)
		}
	} else {
		atomic.StoreInt32(&clt.isValid, isValid)
	}
	return
}

func (clt *client) Write(b []byte) (n int, err error) {
	n, err = clt.conn.Write(b)
	if err != nil {
		if !isTimeoutError(err) {
			atomic.StoreInt32(&clt.isValid, isInvalid)
		}
	} else {
		atomic.StoreInt32(&clt.isValid, isValid)
	}
	return
}

type timeoutErr interface {
	Timeout() bool
}

func isTimeoutError(err error) bool {
	timeoutErr, _ := err.(timeoutErr)
	if timeoutErr == nil {
		return false
	}
	return timeoutErr.Timeout()
}

func borrow() (clt *client, err error) {
	var conn net.Conn
	select {
	case conn = <- pool:
	default:
		conn, err = net.Dial("tcp", "127.0.0.1:18849")
	}
	if err != nil {
		return nil, err
	}
	clt = &client{
		conn: conn,
		inUse: &sync.WaitGroup{},
		isValid: maybeValid,
	}
	return
}

func release(clt *client) error {
	clt.conn.SetDeadline(time.Now().Add(-time.Second))
	clt.inUse.Done()
	clt.inUse.Wait()
	if clt.isValid == isValid {
		return clt.conn.Close()
	}
	select {
	case pool <- clt.conn:
		// returned to pool
		return nil
	default:
		// pool is overflow
		return clt.conn.Close()
	}
}

func handle(server *net.TCPConn) {
	defer server.Close()
	clt, err := borrow()
	if err != nil {
		fmt.Print(err)
		return
	}
	clt.inUse.Add(1)
	defer release(clt)
	go func() {
		clt.inUse.Add(1)
		defer server.Close()
		defer release(clt)
		buf := make([]byte, 2048)
		io.CopyBuffer(server, clt, buf)
	}()
	buf := make([]byte, 2048)
	io.CopyBuffer(clt, server, buf)
}

