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
		description: "simple 1,2,3",
		in:          []int{1, 2, 3},
		expected:    []int{3, 2, 1},
	},
	{
		description: "1 int",
		in:          []int{3},
		expected:    []int{3},
	},
}

func TestReverse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Reverse(tc.in)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for _, tc := range testCases {
		for range b.N {
			Reverse(tc.in)
		}
	}
}
