package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, e := os.ReadFile(filename) //reads whole contents
		if e != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", e)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for k, v := range counts {
		if v > 1 {
			fmt.Printf("\t%s\t%d\n", k, v)
		}
	}
}
