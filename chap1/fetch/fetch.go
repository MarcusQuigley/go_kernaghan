package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, e := http.Get(url)
		if e != nil {
			log.Print(e)
			continue
		}
		defer resp.Body.Close()
		b, e := io.ReadAll(resp.Body)

		if e != nil {
			log.Print("error reading", e)
			continue
		}
		fmt.Printf("%s\n", b)

	}
}
