package p089

func grayCode(n int) []int {
	code := []int {}
	for i := 0; i < 1<<uint(n); i++ {
		code = append(code, i ^ (i>>1))
	}
	return code
}
