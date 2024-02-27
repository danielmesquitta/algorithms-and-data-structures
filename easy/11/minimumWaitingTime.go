package easy11

import (
	"sort"
)

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	currentSum := 0
	totalSum := 0

	for index, query := range queries {
		if index+1 == len(queries) {
			break
		}
		currentSum += query
		totalSum += currentSum
	}

	return totalSum
}
