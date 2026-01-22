package comma

//Benchmarks
//concat - 55.88 ns/op           16 B/op          2 allocs/op
//buffer -
//builder -

// 1234567
// 1234,567

func AddComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return AddComma(s[:n-3]) + "," + s[n-3:]
}
