package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	concat(os.Args[:])
	//join(os.Args[:])
}

func concat(args []string) {
	var s string
	for i := 0; i < len(args); i++ {
		s += fmt.Sprintf("%v. %s\n", i, args[i])
		//sep = " "
	}
	fmt.Println(s)
}

func join(args []string) {
	fmt.Println(strings.Join(args, " "))
}
