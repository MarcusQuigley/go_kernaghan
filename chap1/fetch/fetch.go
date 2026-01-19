package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var datalength int
	for _, url := range os.Args[1:] {
		datalength += fetch(url)
	}

	fmt.Printf("size: %d - elapsed: %v\n", datalength, time.Since(start).Seconds())
}

func fetch(url string) int {

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	response, e := http.Get(url)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error getting: %v", e)
		return 0
	}
	defer response.Body.Close()
	data, e := io.ReadAll(response.Body)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error reading: %v", e)
		return 0
	}
	return len(data)
}
