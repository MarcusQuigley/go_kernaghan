package anagram

import "testing"

var testCases = []struct {
	description string
	ina         string
	inb         string
	expected    bool
}{

	{
		description: "anagram with 3 chars",
		ina:         "ada",
		inb:         "daa",
		expected:    true,
	},
	{
		description: "no anagram with 4 chars",
		ina:         "mike",
		inb:         "kike",
		expected:    false,
	},
	{
		description: "no anagram with diff lens",
		ina:         "gtrhryjr",
		inb:         "31rte1",
		expected:    false,
	},
	{
		description: "no anagram with 4 chars",
		ina:         "eert",
		inb:         "eeer",
		expected:    false,
	},
	{
		description: "no anagram with 5 chars",
		ina:         "eert",
		inb:         "eeer",
		expected:    false,
	},
	{
		description: "no anagram with jap chars",
		ina:         "今天是个好天气，啊好天气",
		inb:         "好天气个啊好天气是，天今",
		expected:    true,
	},

	{"a's", "a", "a", true},
	{"a and b", "a", "b", false},
	{"abc both", "abc", "cba", true},
	{"aabb both", "aabb", "abab", true},
	{"multicase true", "aAbBcC", "abcABC", true},
	{"multicase false", "aAbBcC", "abcABc", false},
}

func TestAHasAnagram(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := HasAnagram(tc.ina, tc.inb)
			if got != tc.expected {
				t.Errorf("got %t. wanted %t", got, tc.expected)
			}
		})
	}
}

func BenchmarkHasAnagram(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				HasAnagram(tc.ina, tc.inb)
			}
		}
	}
}

func TestAHasAnagram2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := HasAnagram2(tc.ina, tc.inb)
			if got != tc.expected {
				t.Errorf("got %t. wanted %t", got, tc.expected)
			}
		})
	}
}

func BenchmarkHasAnagram2(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				HasAnagram2(tc.ina, tc.inb)
			}
		}
	}
}
