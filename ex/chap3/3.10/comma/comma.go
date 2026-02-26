package comma

import (
	"bytes"
	"strings"
)

//Benchmarks
//concat - 55.88 ns/op           16 B/op          2 allocs/op
//buffer - 97.34 ns/op          144 B/op          4 allocs/op
//builder - 68.13 ns/op           32 B/op          3 allocs/op

// 1234567
// 1234,567

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

func AddCommaBuilder(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if i%3 == 1 {
			sb.WriteString(",")
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func AddComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	d, t := len(s)%3, len(s)/3

	if d != 0 {
		buf.WriteString(s[:d])
		buf.WriteByte(',')
	}

	for i := 0; i < t; i++ {
		buf.WriteString(s[i*3+d : d+i*3+3])
		if i != t-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func AddComma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	//12345
	var buf bytes.Buffer
	for i := range n {
		if (n-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}
