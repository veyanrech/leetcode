package gasstation

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {

	globalfuel := 0
	fuelleft := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		globalfuel += diff
		fuelleft += diff

		if fuelleft < 0 {
			start = i + 1
			fuelleft = 0
		}

	}

	if globalfuel < 0 {
		return -1
	}

	return start
}

func Run() {
	// gas := []int{2, 3, 4}
	// cost := []int{3, 4, 3}

	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}

	gas := []int{3, 1, 1}
	cost := []int{1, 2, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}
