package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var flags = flag.String("mode", "sha256", "hash: 'sha256' | 'sha384' | 'sha512'")

func main() {
	flag.Parse()
	bytes, e := io.ReadAll(os.Stdin)
	if e != nil {
		fmt.Printf("Error %v", e)
		return
	}

	switch *flags {
	case "sha256":
		fmt.Printf("sha256 %v", sha256.Sum256(bytes))
	case "sha384":
		fmt.Printf("sha384 %v", sha512.Sum384(bytes))
	case "sha512":
		fmt.Printf("sha512 %v", sha512.Sum512(bytes))
	default:
		fmt.Printf("'%v' is an invalid flag", *flags)
	}
}
