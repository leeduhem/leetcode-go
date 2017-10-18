package p044_1

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

	in00  := "abefcdgiescdfimde"
	in01  := "ab*cd?i*de"
	out00 := true
	testCase(in00, in01, out00, t)

	in10  := "a"
	in11  := "aa"
	out10 := false
	testCase(in10, in11, out10, t)

	in20  := "c"
	in21  := "*?*"
	out20 := true
	testCase(in20, in21, out20, t)

	in30  := "aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba"
	in31  := "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*"
	out30 := true
	testCase(in30, in31, out30, t)
}
