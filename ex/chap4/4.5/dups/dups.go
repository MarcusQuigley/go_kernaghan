package dups

func RemoveDups(sl []string) []string {
	if len(sl) > 1 {

		for i := 1; i < len(sl); {
			if sl[i-1] == sl[i] {
				copy(sl[i:], sl[i+1:])
				sl = sl[:len(sl)-1]
			} else {
				i++
			}
		}
	}
	return sl
}
