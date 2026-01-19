package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// ./fetchdups http://gopl.io msn.com meta.com google.com fox.com nike.com
func main() {
	start := time.Now()
	ch := make(chan string)
	if len(os.Args) == 1 {
		return
	}
	for _, arg := range os.Args[1:] {
		go fetch(arg, ch)
	}
	f, e := os.Create(fmt.Sprintf("./testresults/dupslist_%v.txt", time.Now().Format("15:04:05")))
	if e != nil {
		fmt.Println("Cant create file. error:", e)
		return
	}

	for range os.Args[1:] {
		_, e := f.WriteString(<-ch + "\n")
		if e != nil {
			fmt.Println("Cant write to file. error:", e)
			return
		}
	}
	_, e = fmt.Fprintf(f, " elapsed: %.2fs\n", time.Since(start).Seconds())
	if e != nil {
		fmt.Println("Cant write to file. error:", e)
	}

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	response, e := http.Get(url)
	if e != nil {
		ch <- fmt.Sprintf("error getting: %v", e)
		return
	}
	defer response.Body.Close()

	nbytes, e := io.ReadAll(response.Body)
	if e != nil {
		ch <- fmt.Sprintf("error reading: %v", e)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), len(nbytes), url)
}
