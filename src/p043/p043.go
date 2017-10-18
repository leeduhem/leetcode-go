package p043

func reverseInts(a []int) {
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
}

func char2int(c rune) int {
	return int(c-'0')
}

func int2char(n int) rune {
	return rune(n + '0')
}

func str2rdigits(num string) []int {
	digits := make([]int, len(num))

	for i,c := range num {
		digits[i] = char2int(c)
	}

	reverseInts(digits)
	return digits
}

func rdigits2str(digits []int) string {
	reverseInts(digits)

	str := make([]rune, len(digits))
	for i,n := range digits {
		str[i] = int2char(n)
	}

	return string(str)
}

func timeBy(digits []int, n int) []int {
	digits1 := make([]int, len(digits))
	carry := 0

	for i,d := range digits {
		p := d * n + carry
		digits1[i] = p % 10
		carry = p / 10
	}
	if carry != 0 {
		digits1 = append(digits1, carry)
	}

	return digits1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func digitsSum(digits1, digits2 []int) []int {
	len1, len2 := len(digits1), len(digits2)
	digits := make([]int, max(len1, len2))

	i, carry := 0, 0
	for ; i < len1 && i < len2; i++ {
		s := digits1[i] + digits2[i] + carry
		digits[i] = s % 10
		carry = s / 10
	}

	for ; i < len1; i++ {
		s := digits1[i] + carry
		digits[i] = s % 10
		carry = s / 10
	}

	for ; i < len2; i++ {
		s := digits2[i] + carry
		digits[i] = s % 10
		carry = s / 10
	}

	if carry != 0 {
		digits = append(digits, carry)
	}

	return digits
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	digits1 := str2rdigits(num1)
	digits2 := str2rdigits(num2)

	product := []int {}
	for i,d := range digits2 {
		p := timeBy(digits1, d)
		product = digitsSum(product, append(make([]int, i), p...))
	}

	return rdigits2str(product)
}
