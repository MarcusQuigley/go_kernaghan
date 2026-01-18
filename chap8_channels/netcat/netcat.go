package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, e := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 8000,
	})
	// conn, e := tls.Dial("tcp", "localhost:8000", &tls.Config{
	// 	ServerName: "localhost",
	// })
	if e != nil {
		log.Fatal(e)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, e := io.Copy(w, r); e != nil {
		log.Fatal(e)
	}
}
