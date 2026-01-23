package constants

import (
	//"kernaghan/ex/chap3/3.13/constants"
	"testing"
)

var testCases = []struct {
	description string
	in          any
	expected    int
}{

	{
		description: "anagram with 3 chars",
		in:          KB,
		expected:    1000,
	},
	{
		description: "no anagram with 4 chars",
		in:          GB,
		expected:    1000000,
	},
}

func TestConstants(t *testing.T) {
	//for _, tc := range testCases {
	t.Run("KiB returns 1000", func(t *testing.T) {
		got := KB
		expected := 1000
		if got != expected {
			t.Errorf("got %v. wanted %v", got, expected)
		}
	})
	//}
}
