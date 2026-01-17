package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:8000")
	if e != nil {
		log.Fatal(e)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, e := io.Copy(w, r); e != nil {
		log.Fatal(e)
	}
}
