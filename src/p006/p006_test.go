package p006

import "testing"

func testCase(s, r string, n int, t *testing.T) {
	r1 := convert(s, n)
	if (r1 != r) {
		t.Errorf("Expected %v, got %v", r, r1)
	}
}

func TestConvert(t *testing.T) {
	s11 := "PAYPALISHIRING"
	r10 := "PAHNAPLSIIGYIR"
	testCase(s11, r10, 3, t)

	s21 := "ABCDE"
	r20 := "ABCED"
	testCase(s21, r20, 4, t)

	s31 := ""
	r30 := ""
	testCase(s31, r30, 1, t)

	s41 := "AB"
	r40 := "AB"
	testCase(s41, r40, 1, t)

	s51 := "ABCDEF"
	r50 := "ABFCED"
	testCase(s51, r50, 4, t)
}
