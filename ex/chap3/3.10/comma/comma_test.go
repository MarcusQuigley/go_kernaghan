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

func TestAddCommaBuffer(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := AddCommaBuffer(tc.in)
			if got != tc.expected {
				t.Errorf("got %s. wanted %s", got, tc.expected)
			}
		})
	}
}

func TestAddCommaBuilder(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := AddCommaBuilder(tc.in)
			if got != tc.expected {
				t.Errorf("got %s. wanted %s", got, tc.expected)
			}
		})
	}
}

func TestAddComma2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := AddComma2(tc.in)
			if got != tc.expected {
				t.Errorf("got %s wanted %s", got, tc.expected)
			}
		})
	}
}

func BenchmarkAddCommaBuilder(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				AddCommaBuilder(tc.in)
			}
		}
	}
}

func BenchmarkAddCommaBuffer(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				AddCommaBuffer(tc.in)
			}
		}
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

func BenchmarkAddComma2(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		{
			for _, tc := range testCases {
				AddComma2(tc.in)
			}
		}
	}
}
