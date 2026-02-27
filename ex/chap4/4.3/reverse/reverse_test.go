package reverse

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	description string
	in          []int
	expected    []int
}{
	{
		description: "simple 1 2 3",
		in:          []int{1, 2, 3},
		expected:    []int{3, 2, 1},
	},
}

var testCasesPtr = []struct {
	description string
	in          [3]int
	expected    [3]int
}{
	{
		description: "simple 1 2 3",
		in:          [3]int{1, 2, 3},
		expected:    [3]int{3, 2, 1},
	},
}

func TestReversePtr(t *testing.T) {
	for _, tc := range testCasesPtr {
		t.Run(tc.description, func(t *testing.T) {
			ReversePointer(&tc.in)
			if !reflect.DeepEqual(tc.in, tc.expected) {
				t.Errorf("got %d wanted %d", tc.in, tc.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			Reverse(tc.in)
			if !reflect.DeepEqual(tc.in, tc.expected) {
				t.Errorf("got %d wanted %d", tc.in, tc.expected)
			}
		})
	}
}
func BenchmarkReverse(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				Reverse(tc.in)
			}
		}
	}
}
