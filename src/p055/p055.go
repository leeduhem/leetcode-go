package p055

var mem = map[int]bool {}

func dfs(tree []int, n int) bool {
	if n >= len(tree)-1 || n+tree[n] >= len(tree)-1 {
		mem[n] = true
		return true
	}

	if tree[n] == 0 {
		mem[n] = false
		return false
	}

	// Search from right to left to improve performance for special
	// inputs such as [25000, 24999, 24998, ..., 3, 2, 1, 1, 0, 0]
	for i := n + tree[n]; i >= n+1; i-- {
		if v,ok := mem[i]; ok {
			// Already searched
			return v
		}

		if dfs(tree, i) {
			mem[i] = true
			return true
		}
		//mem[i] = false
		//
		// For almost every elements in [25000, 24999, 24998, ..., 3, 2, 1, 1, 0, 0],
		// searching in big memory is slower than checking them directly through
		// recursive function calls.
	}

	mem[n] = false
	return false
}

func canJump(nums []int) bool {
	// clear memory
	mem = map[int]bool {}

	return dfs(nums, 0)
}
