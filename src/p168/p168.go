package p168

func convertToTitle(n int) string {
	if 1 <= n && n <= 26 {
		return string("ABCDEFGHIJKLMNOPQRSTUVWXYZ"[n-1])
	}

	q, r := n / 26, n % 26
	if r == 0 {
		q -= 1
		r += 26
	}
	return convertToTitle(q) + convertToTitle(r)
}
