package anagram

// 1938 ns/op             328 B/op          3 allocs/op
func HasAnagram(s, y string) bool {
	if len(s) != len(y) {
		return false
	}
	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}
	for _, c := range y {
		if _, ok := m[c]; !ok {
			return false
		}
		m[c]--
		if m[c] == 0 {
			delete(m, c)
		}

	}
	return len(m) == 0
}

// 136.3 ns/op             0 B/op          0 allocs/op
func HasAnagram2(s, y string) bool {
	if len(s) != len(y) {
		return false
	}
	var s_total_unicode, y_total_unicode int
	for _, v := range s {
		s_total_unicode += int(v)
	}

	for _, v := range y {
		y_total_unicode += int(v)
	}
	return s_total_unicode == y_total_unicode
}
