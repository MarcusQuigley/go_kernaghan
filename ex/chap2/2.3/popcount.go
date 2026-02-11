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

// 0.3855 ns/op          0 B/op          0 allocs/op
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 4.106 ns/op           0 B/op          0 allocs/op
func PopCountIter(x uint64) int {
	var result byte
	for i := range 8 {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

func main() {
	for _, num := range os.Args[1:] {
		x, err := strconv.ParseUint(num, 10, 64)
		if err != nil {
			log.Printf("popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(PopCount(x))
		//fmt.Println(PopCountIter(x))
	}
}
