package arraydiff

import "crypto/sha256"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Diff(a, b []byte) int {
	x := sha256.Sum256(a)
	y := sha256.Sum256(b)
	return BitsDiff(&x, &y)
}

// 149.4 ns/op             0 B/op          0 allocs/op
// 161.3 ns/op             0 B/op          0 allocs/op
func BitsDiff(a, b *[sha256.Size]byte) int {
	var result int

	for i := range sha256.Size {
		result += int(pc[a[i]] ^ pc[b[i]])
	}
	return result
}
