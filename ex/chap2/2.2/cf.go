package main

import (
	"bufio"
	"fmt"
	"kernaghan/chap2/tempconv"
	conversions "kernaghan/ex/chap2/2.2/tempconv"
	"os"
	"strconv"
)

func uc(t float64) {
	// temperature
	fmt.Println("Temperature Conversion:")
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n\n",
		f, f.FToC(), c, c.CToF())

	// length
	fmt.Println("Length Conversion:")
	lf := conversions.Feet(t)
	lm := conversions.Metre(t)
	fmt.Printf("%s = %s, %s = %s\n\n",
		lf, lf.FToM(), lm, lm.MToF())
}

func main() {
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			uc(t)
		}
	} else {
		for {
			input := bufio.NewReader(os.Stdin)
			fmt.Fprintf(os.Stdout, "=> ")
			s, err := input.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			// strip `\n`
			t, err := strconv.ParseFloat(s[:len(s)-1], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			uc(t)
		}
	}
}
