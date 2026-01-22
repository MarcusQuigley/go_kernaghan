package unicode

import (
	"fmt"
	"unicode/utf8"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 30.29 ns/op            0 B/op          0 allocs/op
func Contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if HasPrefix(s[i:i+len(substr)], substr) {
			return true
		}
	}
	return false
}
func main() {
	s := "Hello, 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
