package conversions

func (f Feet) FToM() Metre {
	return Metre(f * .3048)
}

func (m Metre) MToF() Feet {
	return Feet(m * 3.28083)
}

func (lb Pounds) LBToK() Kilos {
	return Kilos(lb * .3048)
}

func (k Kilos) KToLB() Pounds {
	return Pounds(k * 3.28083)
}
