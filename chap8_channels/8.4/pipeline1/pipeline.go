package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() { //counter
		for x := range 50 {
			naturals <- x
			time.Sleep(200 * time.Millisecond)

		}
		close(naturals)
	}()

	go func() { //squarer
		for x := range naturals { //nice! no need to use <-naturals
			squares <- x * x
		}
		close(squares)
	}()
	for x := range squares {
		fmt.Println(x)
	}

}
