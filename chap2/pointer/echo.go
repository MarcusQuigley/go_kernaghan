package main

import (
	"flag"
	"fmt"
	"kernaghan/chap2/tempconv"
	"strings"
)

var n = flag.Bool("n", false, "trailing newline")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	//tempconv.BoilingC
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
		fmt.Println(tempconv.BoilingC)
	}
}
