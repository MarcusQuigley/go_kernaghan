package main

import "testing"

var testPopCountCases = []struct {
	description string
	in          uint64
	expected    int
}{
	{
		description: "pc[5]=2",
		in:          5,
		expected:    2,
	},
	{
		description: "pc[15]=4",
		in:          15,
		expected:    4,
	},
}

// 0.3852 ns/op          0 B/op          0 allocs/op
func TestPopCount(t *testing.T) {
	for _, tc := range testPopCountCases {
		t.Run(tc.description, func(t *testing.T) {
			got := PopCount(tc.in)
			if got != tc.expected {
				t.Errorf("got %d. wanted %0d", got, tc.expected)
			}
		})
	}
}

// 4.125 ns/op           0 B/op          0 allocs/op
func BenchmarkPopCount(b *testing.B) {
	for i := range b.N {
		PopCount(uint64(i))
	}
}

func TestPopCountIter(t *testing.T) {
	for _, tc := range testPopCountCases {
		t.Run(tc.description, func(t *testing.T) {

			got := PopCountIter(tc.in)
			if got != tc.expected {
				t.Errorf("got %d. wanted %0d", got, tc.expected)
			}
		})
	}
}

func BenchmarkPopCountIter(b *testing.B) {
	for i := range b.N {
		PopCountIter(uint64(i))
	}
}
