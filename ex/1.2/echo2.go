package main

import (
	"fmt"
	"os"
)

func main() {
	concat(os.Args[:])

}
func concat(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}
