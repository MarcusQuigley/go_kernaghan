package comma

import (
	"bytes"
	"strings"
)

//Benchmarks
//concat  - 55.88 ns/op           16 B/op          2 allocs/op
//buffer  - 97.34 ns/op          144 B/op          4 allocs/op
//builder - 68.13 ns/op           32 B/op          3 allocs/op

// buffer - 406.2 ns/op           464 B/op      12 allocs/op
// 1234567
// 1234,567
func AddComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var end string
	var startindex int = 0
	if k := strings.Index(s, "."); k != -1 {
		end = s[k:]
		n = n - (n - k)
	}

	var buf bytes.Buffer
	if strings.HasPrefix(s, "+") ||
		strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		startindex = 1
	}
	for i := startindex; i < n; i++ { // range n {

		if (n-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteByte(s[i])
	}
	buf.WriteString(end)
	return buf.String()
}
