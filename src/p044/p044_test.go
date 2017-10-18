package p044

import "testing"

func testCase(in0, in1 string, out0 bool, t *testing.T) {
	out1 := isMatch(in0, in1)
	if out0 != out1 {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestIsMatch(t *testing.T) {
	in1 := []string {"aa" , "aa", "aaa", "aa", "aa", "ab", "aab"  }
	in2 := []string {"a"  , "aa", "aa" , "*" , "a*", "?*", "c*a*b"}
	out := []bool   {false, true, false, true, true, true, false  }

	for i := 0; i < len(out); i++ {
		testCase(in1[i], in2[i], out[i], t)
	}
}
