package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	listener, e := net.Listen("tcp", "localhost:8484")
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

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")
		if len(cmds) > 0 {
			switch cmds[0] {
			case "ls":
				exeCmd(c, cmds[0], cmds[1:]...)
			case "cd":
				if e := os.Chdir(cmds[1]); e != nil {
					log.Print(e)
				}
			case "get":
				f, e := os.Open(cmds[1])
				if e != nil {
					log.Printf("file %s: %v", cmds[1], e)
					continue
				}
				mustCopy(c, f)
			case "close":
				return
			default:
				help := ""
				mustCopy(c, strings.NewReader(help))
			}
		}
	}
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, e := io.Copy(w, r); e != nil {
		log.Fatal(e)
	}
}

func exeCmd(w io.Writer, e string, args ...string) {
	cmd := exec.Command(e, args...)
	cmd.Stdout = w
	if e := cmd.Run(); e != nil {
		log.Print(e)
	}
}
