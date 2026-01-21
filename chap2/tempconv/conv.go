package tempconv

func (f Fahrenheit) FToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
