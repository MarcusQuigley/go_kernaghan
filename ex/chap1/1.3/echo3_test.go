package echo

import "testing"

func TestConcat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Concat(tc.args)
			if got == tc.expected {
				t.Fatalf("Concat\n got:%#v\nwant:%#v", got, tc.expected)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Join(tc.args)
			if got == tc.expected {
				t.Fatalf("Concat\n got:%#v\nwant:%#v", got, tc.expected)
			}
		})
	}
}

// 428.3 ns/op           232 B/op         11 allocs/op
func BenchmarkConcat(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Concat")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Concat(tc.args)
		}
	}
}

// 42.82 ns/op           32 B/op          1 allocs/op
func BenchmarkJoin(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping BM Join")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Join(tc.args)
		}
	}
}
