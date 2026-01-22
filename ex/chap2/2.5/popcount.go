package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// x&(x-1)
func PopCount(x uint64) int {

	var result int
	for x != 0 {
		x = x & (x - 1) //clears the rig htmost non-zero bit of x
		result++
	}
	return result
}

func main() {
	for _, num := range os.Args[1:] {
		x, err := strconv.ParseUint(num, 10, 64)
		if err != nil {
			log.Printf("popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(PopCount(x))
	}
}
