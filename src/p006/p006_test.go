package p006

import (
	"testing"
)

func TestConvert(t *testing.T) {
	var tests = []struct {
		in0  string
		in1  int
		want string
	}{
		{"", 1, ""},
		{"AB", 1, "AB"},
		{"ABCDE", 4, "ABCED"},
		{"ABCDEF", 4, "ABFCED"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	}

	for _, test := range tests {
		got := convert(test.in0, test.in1)
		if got != test.want {
			t.Errorf("convert(%q, %v) = %q, want %q",
				test.in0, test.in1, got, test.want)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	in0, in1, want := "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"
	for i := 0; i < b.N; i++ {
		got := convert(in0, in1)
		if got != want {
			b.Errorf("convert(%q, %v) = %q, want %q",
				in0, in1, got, want)
		}
	}
}
