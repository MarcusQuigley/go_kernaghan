package main

import (
	"bytes"
	"fmt"
)

func main() {
	musings()
}

func musings() {
	var sym1 = "¿"
	var sym = '¿'

	fmt.Printf("%x %v\n", sym1, sym1)
	fmt.Printf("%x %v\n", sym, sym)
	s := "abc¿"
	by := []byte(s)
	fmt.Println(s)
	fmt.Println(by)
	var buf bytes.Buffer
	fmt.Println(buf.String())
	buf.WriteRune('¿')
	fmt.Println(buf.String())
	buf.WriteString("sh1")
	fmt.Println(buf.String())

}
