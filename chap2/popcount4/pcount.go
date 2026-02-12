package main

var pc [16]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// [0,1,1,2,
//  1,2,2,3,
//  1,2,2,3,
//  2,3,3,4]

func PopCount(x uint64) int {
	var result byte
	for i := range 4 {
		result += pc[byte(x>>(i*4))]
	}
	return int(result)
}
