package p005

func isPalindrome(s string) bool {
	l := len(s)
	h := l / 2

	for i,v := range s[0:h] {
		//if (v != []rune(s)[l-1-i]) {
		// is more proper, but leetcode thought it's too slow.
		if (byte(v) != s[l-1-i]) {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {
	l   := len(s)
	low,high := 0,0
	max := 0

	for i,_ := range s {
		for j := i+1; j <= l; j++ {
			if (isPalindrome(s[i:j]) && (j-i+1 > max)) {
				max = j-i+1
				low,high = i,j
			}
		}
	}

	return s[low:high]
}
