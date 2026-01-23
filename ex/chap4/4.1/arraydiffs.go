package arraydiff

import "crypto/sha256"

// 149.4 ns/op             0 B/op          0 allocs/op
func Diffs(a, b []byte) int {
	var result int
	ash := sha256.Sum256(a)
	bsh := sha256.Sum256(b)

	for i := range ash {
		if ash[i] != bsh[i] {
			result++
		}
	}
	return result
}
