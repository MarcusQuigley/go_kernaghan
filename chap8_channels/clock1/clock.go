package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, e := net.Listen("tcp", "localhost:8880")
	if e != nil {
		log.Fatal(e)
	}
	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Print(e)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, e := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if e != nil {
			return //conn dead
		}
		time.Sleep(1 * time.Second)
	}
}
