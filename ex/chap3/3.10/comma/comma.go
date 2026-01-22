package comma

import "bytes"

//Benchmarks
//concat - 55.88 ns/op           16 B/op          2 allocs/op
//buffer - 97.34 ns/op          144 B/op          4 allocs/op
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

func AddCommaBuffer(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := range n {

		if (n-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}
