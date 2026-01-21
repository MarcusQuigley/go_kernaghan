package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, e := http.Get(url)
		if e != nil {
			log.Printf("Get %v\n", e)
			continue
		}
		defer resp.Body.Close()

		fmt.Fprintf(os.Stdout, "Status: %s", resp.Status)
		_, e = io.Copy(os.Stdout, resp.Body)
		if e != nil {
			log.Printf("Get %v\n", e)
			continue
		}
	}
}
