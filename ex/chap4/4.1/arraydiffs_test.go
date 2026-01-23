package arraydiff

import "testing"

var testCases = []struct {
	description string
	arg1        []byte
	arg2        []byte
	expected    int
}{

	{
		description: "x and X",
		arg1:        []byte("x"),
		arg2:        []byte("X"),
		expected:    31,
	},
}

func TestDiffs(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Diffs(tc.arg1, tc.arg2)
			if got != tc.expected {
				t.Errorf("got %d. wanted %d", got, tc.expected)
			}
		})
	}
}

func BenchmarkDiffs(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				Diffs(tc.arg1, tc.arg2)
			}
		}
	}
}
