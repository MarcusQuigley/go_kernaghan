package comma

import "testing"

var testCases = []struct {
	description string
	in          string
	expected    string
}{

	{
		description: "no commas",
		in:          "123",
		expected:    "123",
	},
	{
		description: "1 comma",
		in:          "1234",
		expected:    "1,234",
	},
	{
		description: "2 commas",
		in:          "1234567",
		expected:    "1,234,567",
	},
	{
		description: "float  and 3 places",
		in:          "1234567.666",
		expected:    "1,234,567.666",
	},
	{
		description: "2 commas float and 1 place",
		in:          "1234567.5",
		expected:    "1,234,567.5",
	},
	{
		description: "2 commas float with positive and 3 places",
		in:          "+1234567.666",
		expected:    "+1,234,567.666",
	},
	{
		description: "2 commas with positive",
		in:          "+1234567",
		expected:    "+1,234,567",
	},
	{
		description: "no commas float and 1 place",
		in:          "3.5",
		expected:    "3.5",
	},
	{
		description: "2 commas float with negative and 3 places",
		in:          "-1234567.666",
		expected:    "-1,234,567.666",
	},
	{
		description: "2 commas with negative",
		in:          "-1234567",
		expected:    "-1,234,567",
	},
}

var testCases2 = []struct {
	description string
	in          string
	expected    string
}{

	{
		description: "2 commas with positive",
		in:          "+1234567",
		expected:    "+1,234,567",
	},
}

func TestAddComma(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := AddComma(tc.in)
			if got != tc.expected {
				t.Errorf("got %s. wanted %s", got, tc.expected)
			}
		})
	}
}

func BenchmarkAddComma(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				AddComma(tc.in)
			}
		}
	}
}

func TestAddComma1(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := AddComma1(tc.in)
			if got != tc.expected {
				t.Errorf("got %s. wanted %s", got, tc.expected)
			}
		})
	}
}

func BenchmarkAddComma1(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				AddComma1(tc.in)
			}
		}
	}
}
