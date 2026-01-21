package tempconv

import "testing"

var testCelsiusCases = []struct {
	description string
	in          Celsius
	expected    Fahrenheit
}{
	{
		description: "c to f good",
		in:          15,
		expected:    59,
	},
}

func TestCtoF(t *testing.T) {
	for _, tc := range testCelsiusCases {
		t.Run(tc.description, func(t *testing.T) {
			var c Celsius = tc.in
			got := c.CToF()
			if got != tc.expected {
				t.Errorf("got %v. wanted %v", got, tc.expected)
			}
		})
	}

}

var testFahrenheitCases = []struct {
	description    string
	in             Fahrenheit
	expected       Celsius
	expectedString string
}{
	{
		description:    "c to f good",
		in:             15,
		expected:       -9.444444444444445,
		expectedString: "-9°C",
	},
}

func TestFtoC(t *testing.T) {
	for _, tc := range testFahrenheitCases {
		t.Run(tc.description, func(t *testing.T) {
			var f Fahrenheit = tc.in
			got := f.FToC()

			if got != tc.expected {
				t.Errorf("got %0.1f. wanted %0.1f", got, tc.expected)
			}
		})
	}

}

// got    -9.444444°C.
// wanted -9.444444°C

func TestFtoCString(t *testing.T) {
	for _, tc := range testFahrenheitCases {
		t.Run(tc.description, func(t *testing.T) {
			var f Fahrenheit = tc.in
			got := f.FToC()

			if got.String() != tc.expectedString {
				t.Errorf("got %v. wanted %v", got, tc.expectedString)
			}
		})
	}

}
