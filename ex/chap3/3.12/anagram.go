package anagram

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
