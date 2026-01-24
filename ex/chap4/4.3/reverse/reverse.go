package reverse

// ptr *[32]byte)
// 3.617 ns/op           0 B/op          0 allocs/op
func Reverse(in *[3]int) *[3]int {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}
