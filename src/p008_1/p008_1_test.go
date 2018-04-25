package p008_1

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	var tests = []struct {
		in   string
		want int
	}{
		{"", 0}, {"+", 0},
		{"10", 10},
		{" 123", 123}, {"  -123", -123},
		{"123 ", 123},
		{"2147483648", 2147483647},
		{"9223372036854775809", 2147483647},
		{"-6147483648", -2147483648},
		{"-9223372036854775809", -2147483648},
	}

	for _, test := range tests {
		got := myAtoi(test.in)
		if got != test.want {
			t.Errorf("myAtoi(%q) = %v, want %v",
				test.in, got, test.want)
		}
	}

}
