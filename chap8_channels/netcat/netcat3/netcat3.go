package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:8083")
	if e != nil {
		log.Fatal(e)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	if c, ok := conn.(*net.TCPConn); ok {
		c.CloseWrite()
	} else {
		conn.Close()
	}
	<-done //wait for goro to finish
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, e := io.Copy(w, r); e != nil {
		log.Fatal(e)
	}
}
