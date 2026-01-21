package main

import "fmt"

type Celcius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
	FreezingK     Kelvin  = -273.15
)

func CtoF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)

}

func FtoC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func CtoK(c Celcius) Kelvin {
	return Kelvin(c + Celcius(FreezingK))
}
func (k Fahrenheit) String() string { return fmt.Sprintf("%g°F", k) }
func (k Celcius) String() string    { return fmt.Sprintf("%g°C", k) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
