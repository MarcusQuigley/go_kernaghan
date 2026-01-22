package unicode

import "testing"

var testCases = []struct {
	description string
	in          string
	substr      string
	expected    bool
}{
	{
		description: "is he in hello",
		in:          "hello",
		substr:      "he",
		expected:    true,
	},
	{
		description: "is lo in hello",
		in:          "hello",
		substr:      "lo",
		expected:    true,
	},
	{
		description: "is lo in hello",
		in:          "hello",
		substr:      "olo",
		expected:    false,
	},
}

func TestContains(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Contains(tc.in, tc.substr)
			if got != tc.expected {
				t.Errorf("got %t. wanted %t", got, tc.expected)
			}
		})
	}
}

func BenchmarkContains(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				Contains(tc.in, tc.substr)
			}
		}
	}
}
