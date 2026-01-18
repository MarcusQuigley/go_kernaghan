package echo

var testCases = []struct {
	description string
	args        []string
	expected    string
}{
	{
		description: "few names",
		args:        []string{"hello", "world", "zombies", "pants"},
		expected:    "",
	},
}
