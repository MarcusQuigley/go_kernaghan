package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)

	for _, filename := range os.Args[1:] {
		data, e := os.ReadFile(filename)
		if e != nil {
			fmt.Fprintf(os.Stderr, "%v\n", e)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line] = append(counts[line], filename)
		}

	}
	for _, v := range counts {
		if len(v) > 1 {
			fmt.Fprintf(os.Stderr, "\t%v\n", v)
		}
	}
}
