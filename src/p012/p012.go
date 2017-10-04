package p012

import "strings"

func baseInt2Roman(n int) string {
	switch {
	case n % 1000 == 0:
		return strings.Repeat("M", n / 1000)
	case n % 100 == 0:
		return strings.Repeat("C", n / 100)
	case n % 10 == 0:
		return strings.Repeat("X", n / 10)
	default:
		return strings.Repeat("I", n)
	}
}

func intToRoman(num int) string {
	if num < 0 || num > 3999 {
		panic("Valid input range: 0 ~ 3999")
	}

	switch {
	case num >= 1000:
		d,r := num / 1000, num % 1000
		return baseInt2Roman(d * 1000) + intToRoman(r)
	case num >= 900:
		return "CM" + intToRoman(num - 900)
	case num >= 500:
		return "D"  + intToRoman(num - 500)
	case num >= 400:
		return "CD" + intToRoman(num - 400)
	case num >= 100:
		d,r := num / 100, num % 100
		return baseInt2Roman(d * 100) + intToRoman(r)
	case num >= 90:
		return "XC" + intToRoman(num - 90)
	case num >= 50:
		return "L"  + intToRoman(num - 50)
	case num >= 40:
		return "XL" + intToRoman(num - 40)
	case num >= 10:
		d,r := num / 10, num % 10
		return baseInt2Roman(d * 10) + intToRoman(r)
	case num >= 9:
		return "IX" + intToRoman(num - 9)
	case num >= 5:
		return "V"  + intToRoman(num - 5)
	case num == 4:
		return "IV"
	case num >= 1:
		return baseInt2Roman(num)
	}

	return ""
}
