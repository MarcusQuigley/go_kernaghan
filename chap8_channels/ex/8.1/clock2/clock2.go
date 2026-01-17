package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8770", "port, number")

func main() {
	listen, e := net.Listen("tcp", "localhost:"+*port)
	if e != nil {
		log.Fatal(e)
	}
	//!+
	for {
		conn, e := listen.Accept()
		if e != nil {
			log.Print(e)
			continue
		}
		go handleConn(conn)
	}
	//!-
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, e := io.WriteString(conn, time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006\n"))
		//_, e := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if e != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
