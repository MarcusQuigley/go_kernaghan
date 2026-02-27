package main

import (
	"fmt"
)

func zero(arr *[32]byte) {
	for i := range arr {
		arr[i] = 0
	}
}

func zero1(arr *[32]byte) {
	*arr = [32]byte{}
}

func nonempty(sl []string) []string {
	var i int
	for _, v := range sl {
		if v != "" {
			sl[i] = v
			i++
		}
	}
	return sl[:i]
}

func check_nonempty() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func slicesmess() {
	mainsl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sub_mainsl2 := mainsl[:4]

	sub_mainsl := mainsl[4:]
	sub_mainsl3 := mainsl[6:10]

	fmt.Printf("len:%d\tcap:%d\t %v\n", cap(mainsl), len(mainsl), mainsl)
	fmt.Printf("len:%d\tcap:%d\t %v\n", cap(sub_mainsl2), len(sub_mainsl2), sub_mainsl2)
	fmt.Printf("len:%d\tcap:%d\t %v\n", cap(sub_mainsl), len(sub_mainsl), sub_mainsl)
	fmt.Printf("len:%d\tcap:%d\t %v\n", cap(sub_mainsl3), len(sub_mainsl3), sub_mainsl3)
	fmt.Println("--------")
	x := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x, " len:", len(x), " cap:", cap(x))
	fmt.Println("y:", y, " len:", len(y), " cap:", cap(y))
	fmt.Println("z:", z, " len:", len(z), " cap:", cap(z))
	fmt.Println("d:", d, " len:", len(d), " cap:", cap(d))
	fmt.Println("e:", e, " len:", len(e), " cap:", cap(e))

	fmt.Println("--------")

	xx := []string{"a", "b", "c", "d"}
	yy := xx[:2]
	zz := xx[1:]
	fmt.Println("xx:", xx)
	fmt.Println("yy:", yy)
	fmt.Println("zz:", zz)
}

func doublelices(sl []int) { //[]int {
	for i := range sl {
		sl[i] *= 2
	}
	//return sl
}

func doublearray(arr *[3]int) {
	for i := range arr {
		(*arr)[i] *= 2
	}
}

func main() {
	sli := []int{1, 2, 3}
	fmt.Println(sli)
	//slxx := doublelices(sli)
	doublelices(sli)
	fmt.Println(sli)
	ari := [3]int{10, 20, 30}
	fmt.Println(ari)
	doublearray(&ari)
	fmt.Println(ari)
	//slicesmess()
	//check_nonempty()
}

// len:10	cap:10	 [1 2 3 4 5 6 7 8 9 10]
// len:8	cap:2	 [3 4]
// len:4	cap:4	 [7 8 9 10]
// len:6	cap:3	 [5 6 7]

// x: [a b c d]  len: 4  cap: 4
// y: [a b]  len: 2  cap: 4
// z: [b c d]  len: 3  cap: 3
// d: [b c]  len: 2  cap: 3
// e: [a b c d]  len: 4  cap: 4
