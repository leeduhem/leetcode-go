package p118

func generateRow(prev []int, n int) []int {
	row := make([]int, n)
	row[0], row[n-1] = 1, 1

	for i := 1; i < n-1; i++ {
		row[i] = prev[i-1] + prev[i]
	}

	return row
}

func generate(numRows int) [][]int {
	rows := make([][]int, numRows+1)

	for i := 1; i <= numRows; i++ {
		rows[i] = generateRow(rows[i-1], i)
	}

	return rows[1:]
}
