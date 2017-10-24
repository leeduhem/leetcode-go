package p070

var mem map[int]int

func climbStairs1(n int) int {
	if v,ok := mem[n]; ok {
		return v
	}

	mem[n] = climbStairs1(n-1) + climbStairs1(n-2)
	return mem[n]
}

func climbStairs(n int) int {
	// clear memory
	mem = map[int]int {0:0, 1:1, 2:2}

	return climbStairs1(n)
}
