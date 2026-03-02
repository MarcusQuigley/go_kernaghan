package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	count()
}

func count() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := strings.NewReader(os.Args[1])
	for {
		r, n, e := in.ReadRune()
		if e == io.EOF {
			break
		}
		if e != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", e)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for k, v := range counts {
		fmt.Printf("%q\t%v\n", k, v)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, v := range utflen {
		fmt.Printf("%q\t%v\n", i, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid utf-8 chars\n", invalid)
	}
}
