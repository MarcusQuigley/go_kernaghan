package dups

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	description string
	in          []string
	expected    []string
}{
	{
		"no dups",
		[]string{"1", "2"},
		[]string{"1", "2"},
	},
	{
		"some dups",
		[]string{"1", "2", "2", "3"},
		[]string{"1", "2", "3"},
	},
	{
		"some dups2",
		[]string{"1", "2", "2", "3", "3"},
		[]string{"1", "2", "3"},
	},
}

func TestRemoveDups(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := RemoveDups(tc.in)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v want %v", got, tc.expected)
			}
		})
	}
}
