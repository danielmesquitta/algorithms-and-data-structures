package medium28

import (
	"math"

	"golang.org/x/exp/constraints"
)

func MaxSubsetSumNoAdjacent(array []int) int {
	switch len(array) {
	case 0:
		return 0
	case 1:
		return array[0]
	}

	maxSums := make([]int, len(array))
	maxSums[0], maxSums[1] = array[0], Max(array[0], array[1])

	for i := 2; i < len(array); i++ {
		maxSums[i] = Max(maxSums[i-1], array[i]+maxSums[i-2])
	}

	return maxSums[len(array)-1]
}

func Max[T constraints.Float | constraints.Integer](a, b T) T {
	return any(math.Max(float64(a), float64(b))).(T)
}
