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

// 0.3837 ns/op           0 B/op          0 allocs/op
//
//	func PopCount(x uint64) int {
//		var result int
//		for range 64 {
//			if x&1 == 1 {
//				result++
//			}
//			x >>= 1
//		}
//		return result
//	}

// 26.70 ns/op            0 B/op          0 allocs/op
func PopCount(x uint64) int {
	var result int
	for range 64 {
		if x == 0 {
			break
		}
		if x&1 == 1 {
			result++
		}
		x >>= 1
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
