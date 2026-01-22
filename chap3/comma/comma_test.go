package comma

import "testing"

var testCases = []struct {
	description string
	in          string
	expected    string
}{
	{
		description: "is he in hello",
		in:          "1234567",
		expected:    "1,234,567",
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
