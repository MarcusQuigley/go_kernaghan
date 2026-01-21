package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// parameter: NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
func main() {
	for _, v := range os.Args[2:] {
		s := strings.Split(v, "=")[1]
		go getTime(s)
	}
	s := strings.Split(os.Args[1], "=")
	getTime(s[1])
}

func getTime(addr string) {
	conn, e := net.Dial("tcp", addr)
	if e != nil {
		log.Print(e)
	}
	mustCopy(conn)
}

func mustCopy(conn net.Conn) {
	if _, e := io.Copy(os.Stdout, conn); e != nil {
		log.Print("when copying.. ", e)
	}
}
