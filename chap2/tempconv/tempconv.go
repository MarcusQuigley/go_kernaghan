package tempconv

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func (c Celcius) String() string {
	return fmt.Sprintf("%0.0f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%0g°F", f)
}

func (f Fahrenheit) FToC() Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func (c Celcius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
