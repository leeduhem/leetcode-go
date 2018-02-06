package p134

func checkIndex(idx int, gas, cost []int) bool {
	tank := 0

	for i := idx; i < len(gas); i++ {
		if tank + gas[i] < cost[i] {
			return false
		}
		tank += gas[i] - cost[i]
	}

	for i := 0; i < idx; i++ {
		if tank + gas[i] < cost[i] {
			return false
		}
		tank += gas[i] - cost[i]
	}

	return true
}

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if checkIndex(i, gas, cost) {
			return i
		}
	}

	return -1
}
