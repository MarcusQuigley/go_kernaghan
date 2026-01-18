package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, e := http.Get(url)
		if e != nil {
			log.Printf("Get %v\n", e)
			continue
		}
		defer resp.Body.Close()
		_, e = io.Copy(os.Stdout, resp.Body)
		if e != nil {
			log.Printf("Get %v\n", e)
			continue
		}
	}
}
