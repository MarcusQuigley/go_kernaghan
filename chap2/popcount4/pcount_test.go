package main

import "testing"

var testPopCases = []struct {
	description string
	in          uint64
	expected    int
}{
	{
		description: "0 is 0",
		in:          0,
		expected:    0,
	},
	{
		description: "3 is 2",
		in:          3,
		expected:    2,
	},
	{
		description: "6 is 6",
		in:          6,
		expected:    2,
	},
	{
		description: "7 is 3",
		in:          7,
		expected:    3,
	},
	{
		description: "15 is 4",
		in:          15,
		expected:    4,
	},
}

func TestPopCount(t *testing.T) {
	for _, tc := range testPopCases {
		t.Run(tc.description, func(t *testing.T) {
			got := PopCount(tc.in)
			if got != tc.expected {
				t.Errorf("got %d wanted %d", got, tc.expected)
			}
		})
	}

}
