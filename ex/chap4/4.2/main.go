package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		fmt.Printf("%v\n", sha256.Sum256([]byte(t)))
	}
}
