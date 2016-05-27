package main

import (
	"fmt"
)

func ReadFull(r Reader, buf []byte) (int, error)  {
	var n int
	var err error

	for len (buf) > 0 {
	 var nr int
	 nr, err = r.Read(buf)
	 n += nr

	 if err != nil {
			return n, err
	 }
	 buf = buf[nr:]
	}
	return n, err
}


func  ReadFullImpro(r Reader, buf, []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
}







