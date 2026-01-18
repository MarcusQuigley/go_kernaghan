package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, e := os.Open(file)
			if e != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", e)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for k, v := range counts {
		fmt.Printf("%d\\t%s\n", v, k)
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
