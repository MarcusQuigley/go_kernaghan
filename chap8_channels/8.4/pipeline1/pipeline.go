package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for x := range 50 {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
	println("Done!")
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	//printer := make(chan<- int)

	go counter(naturals)
	go squarer(squares, naturals)

	// go func() { //counter
	// 	for x := range 50 {
	// 		naturals <- x
	// 		time.Sleep(200 * time.Millisecond)

	// 	}
	// 	close(naturals)
	// }()

	// go func() { //squarer
	// 	for x := range naturals { //nice! no need to use <-naturals
	// 		squares <- x * x
	// 	}
	// 	close(squares)
	// }()
	// for x := range squares {
	// 	fmt.Println(x)
	// }
	printer(squares)

}
