package p093

import (
	"testing"
	. "testing/utils"
)

func testCase(in string, want []string, t *testing.T) {
	got := restoreIpAddresses(in)
	if ! IsStrsEqualWithoutOrder(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestRestoreIpAddresses(t *testing.T) {
	in0 := "25525511135"
	want0 := []string {"255.255.11.135", "255.255.111.35"}
	testCase(in0, want0, t)

	in1 := "010010"
	want1 := []string {"0.10.0.10","0.100.1.0"}
	testCase(in1, want1, t)
}
