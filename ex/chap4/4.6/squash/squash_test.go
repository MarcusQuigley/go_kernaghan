package squash

// {"Hello　世界", "Hello 世界"},
var testCases = []struct {
	description string
	in          string
	expected    string
}{
	{
		"remove 1 space",
		"Hello　 世界",
		"Hello 世界",
	},
	{
		" more",
		"AB C  D   ",
		"AB C D ",
	},
}

// func TestRemoveDups(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run(tc.description, func(t *testing.T) {
// 			got := string(uniSpace([]byte(tc.in)))
// 			if !reflect.DeepEqual(got, tc.expected) {
// 				t.Errorf("got %v want %v", got, tc.expected)
// 			}
// 		})
// 	}
// }
