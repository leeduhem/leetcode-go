package p005_1

import "testing"

func checkResult(r string, rs []string, t *testing.T) {
	ok := false

	for _,r1 := range rs {
		if (r == r1) {
			ok = true
		}
	}

	if !ok {
		t.Errorf("Expected one of %v, got %v", rs, r)
	}
}

func TestLongestPalindrome(t *testing.T) {
	s1  := "babad"
	rs1 := []string {"bab", "aba"}
	r1  := longestPalindrome(s1)
	checkResult(r1, rs1, t)

	s2  := "cbbd"
	rs2 := []string {"bb"}
	r2  := longestPalindrome(s2)
	checkResult(r2, rs2, t)

	s3  := "a"
	rs3 := []string {"a"}
	r3  := longestPalindrome(s3)
	checkResult(r3, rs3, t)

	s4  := ""
	rs4 := []string {""}
	r4  := longestPalindrome(s4)
	checkResult(r4, rs4, t)

	s5  := "bb"
	rs5 := []string {"bb"}
	r5  := longestPalindrome(s5)
	checkResult(r5, rs5, t)

	s6  := "kyyrjtdplseovzwjkykrjwhxquwxsfsorjiumvxjhjmgeueafubtonhlerrgsgohfosqssmizcuqryqomsipovhhodpfyudtusjhonlqabhxfahfcjqxyckycstcqwxvicwkjeuboerkmjshfgiglceycmycadpnvoeaurqatesivajoqdilynbcihnidbizwkuaoegmytopzdmvvoewvhebqzskseeubnretjgnmyjwwgcooytfojeuzcuyhsznbcaiqpwcyusyyywqmmvqzvvceylnuwcbxybhqpvjumzomnabrjgcfaabqmiotlfojnyuolostmtacbwmwlqdfkbfikusuqtupdwdrjwqmuudbcvtpieiwteqbeyfyqejglmxofdjksqmzeugwvuniaxdrunyunnqpbnfbgqemvamaxuhjbyzqmhalrprhnindrkbopwbwsjeqrmyqipnqvjqzpjalqyfvaavyhytetllzupxjwozdfpmjhjlrnitnjgapzrakcqahaqetwllaaiadalmxgvpawqpgecojxfvcgxsbrldktufdrogkogbltcezflyctklpqrjymqzyzmtlssnavzcquytcskcnjzzrytsvawkavzboncxlhqfiofuohehaygxidxsofhmhzygklliovnwqbwwiiyarxtoihvjkdrzqsnmhdtdlpckuayhtfyirnhkrhbrwkdymjrjklonyggqnxhfvtkqxoicakzsxmgczpwhpkzcntkcwhkdkxvfnjbvjjoumczjyvdgkfukfuldolqnauvoyhoheoqvpwoisniv"
	rs6 := []string {"qahaq"}
	r6  := longestPalindrome(s6)
	checkResult(r6, rs6, t)
}
