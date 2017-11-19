package p119

import (
	"math/big"
)

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)

	var n int64 = int64(rowIndex)
	for k := 0; k <= rowIndex; k++ {
		c := big.NewInt(0).Binomial(n, int64(k))
		if ! c.IsInt64() {
			panic("Invalid input")
		}
		row[k] = int(c.Int64())
	}

	return row
}
