package p005

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		in    string
		wants []string
	}{
		{"", []string{""}},
		{"a", []string{"a"}},
		{"cbbd", []string{"bb"}},
		{"babad", []string{"bab", "aba"}},
		{"bb", []string{"bb"}},
		{"世界界世", []string{"世界界世"}},
	}

nextTest:
	for _, test := range tests {
		got := longestPalindrome(test.in)
		for _, want := range test.wants {
			if got == want {
				continue nextTest
			}
		}
		t.Errorf("longestPalindrome(%q) = %q, want one of %#v",
			test.in, got, test.wants)
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	in := "kyyrjtdplseovzwjkykrjwhxquwxsfsorjiumvxjhjmgeueafubtonhlerrgsgohfosqssmizcuqryqomsipovhhodpfyudtusjhonlqabhxfahfcjqxyckycstcqwxvicwkjeuboerkmjshfgiglceycmycadpnvoeaurqatesivajoqdilynbcihnidbizwkuaoegmytopzdmvvoewvhebqzskseeubnretjgnmyjwwgcooytfojeuzcuyhsznbcaiqpwcyusyyywqmmvqzvvceylnuwcbxybhqpvjumzomnabrjgcfaabqmiotlfojnyuolostmtacbwmwlqdfkbfikusuqtupdwdrjwqmuudbcvtpieiwteqbeyfyqejglmxofdjksqmzeugwvuniaxdrunyunnqpbnfbgqemvamaxuhjbyzqmhalrprhnindrkbopwbwsjeqrmyqipnqvjqzpjalqyfvaavyhytetllzupxjwozdfpmjhjlrnitnjgapzrakcqahaqetwllaaiadalmxgvpawqpgecojxfvcgxsbrldktufdrogkogbltcezflyctklpqrjymqzyzmtlssnavzcquytcskcnjzzrytsvawkavzboncxlhqfiofuohehaygxidxsofhmhzygklliovnwqbwwiiyarxtoihvjkdrzqsnmhdtdlpckuayhtfyirnhkrhbrwkdymjrjklonyggqnxhfvtkqxoicakzsxmgczpwhpkzcntkcwhkdkxvfnjbvjjoumczjyvdgkfukfuldolqnauvoyhoheoqvpwoisniv"
	want := "qahaq"

	for i := 0; i < b.N; i++ {
		got := longestPalindrome(in)
		if got != want {
			b.Errorf("longestPalindrome(%q) = %q, want %q",
				in, got, want)
		}
	}
}
