package p200

func markPos(grid [][]byte, r, c int) {
	grid[r][c] = 'x'

	if r-1 >= 0 && grid[r-1][c] == '1' {
		markPos(grid, r-1, c)
	}
	if r+1 < len(grid) && grid[r+1][c] == '1' {
		markPos(grid, r+1, c)
	}
	if c-1 >= 0 && grid[r][c-1] == '1' {
		markPos(grid, r, c-1)
	}
	if c+1 < len(grid[r]) && grid[r][c+1] == '1' {
		markPos(grid, r, c+1)
	}
}

func numIslands(grid [][]byte) int {
	num := 0

	for i, row := range grid {
		for j, col := range row {
			if col == '1' {
				num++
				markPos(grid, i, j)
			}
		}
	}

	return num
}
