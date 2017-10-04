package p012

import "testing"

func testCase(n int, r0 string, t *testing.T) {
	r1 := intToRoman(n)
	if (r0 != r1) {
		t.Errorf("Case %v: expected %v, got %v", n, r0, r1)
	}
}

func TestIntToRoman(t *testing.T) {
	testCase(1, "I", t)
	testCase(2, "II", t)
	testCase(3, "III", t)
	testCase(4, "IV", t)
	testCase(5, "V", t)
	testCase(6, "VI", t)
	testCase(7, "VII", t)
	testCase(8, "VIII", t)
	testCase(9, "IX", t)
	testCase(10, "X", t)
	testCase(40, "XL", t)
	testCase(50, "L", t)
	testCase(90, "XC", t)
	testCase(100, "C", t)
	testCase(400, "CD", t)
	testCase(500, "D", t)
	testCase(900, "CM", t)
	testCase(1000, "M", t)
}
