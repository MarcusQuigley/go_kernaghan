package main

func PopCount(x uint64) int {

	var result int
	for x != 0 {
		x &= (x - 1) //clears the rig htmost non-zero bit of x
		result++
	}
	return result
}
