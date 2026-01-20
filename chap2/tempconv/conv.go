package tempconv

func (f Fahrenheit) FToC() Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func (c Celcius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
