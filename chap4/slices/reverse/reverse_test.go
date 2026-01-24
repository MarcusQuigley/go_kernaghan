package reverse

import (
	"reflect"
	"testing"
)

var testCasesReverse = []struct {
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

var testCasesEqual = []struct {
	description string
	x           []int
	y           []int
	expected    bool
}{
	{
		description: "equal 3 ints",
		x:           []int{1, 2, 3},
		y:           []int{1, 2, 3},
		expected:    true,
	},
	{

		description: "not equal 3 ints",
		x:           []int{1, 2, 3},
		y:           []int{3, 2, 1},
		expected:    false,
	},
	{
		description: "not equal many ints",
		x:           []int{1, 2, 3, 6, 2, 3, 57},
		y:           []int{2, 3, 6, 2, 3, 57, 3},
		expected:    false,
	},
}

func TestReverse(t *testing.T) {
	for _, tc := range testCasesReverse {
		t.Run(tc.description, func(t *testing.T) {
			got := Reverse(tc.in)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	for _, tc := range testCasesEqual {
		t.Run(tc.description, func(t *testing.T) {
			got := Equal(tc.x, tc.y)
			if got != tc.expected {
				t.Errorf("got %t, want %t", got, tc.expected)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for _, tc := range testCasesReverse {
		for range b.N {
			Reverse(tc.in)
		}
	}
}
