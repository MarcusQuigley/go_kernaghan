package conversions

import (
	"fmt"
)

// type Celsius float64
// type Fahrenheit float64
// type Kelvin float64
type Pounds float64
type Kilos float64
type Feet float64
type Metre float64

// const (
// 	// AbsoluteZeroC Celsius = -273.15
// 	// FreezingC     Celsius = 0
// 	// BoilingC      Celsius = 100
// 	AbsoluteZeroK Kelvin = -273.15
// )

// func (k Kelvin) String() string {
// 	return fmt.Sprintf("%0.0f˚K", k)
// }

func (p Pounds) String() string {
	return fmt.Sprintf("%0.0f lb", p)
}

func (k Kilos) String() string {
	return fmt.Sprintf("%0.0f kg", k)
}

func (m Metre) String() string {
	return fmt.Sprintf("%0.0f(m)", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%0.0f(ft)", f)
}
