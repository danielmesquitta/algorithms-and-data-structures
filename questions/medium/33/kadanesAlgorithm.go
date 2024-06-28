package medium33

import (
	"math"
)

func KadanesAlgorithm(array []int) int {
	greatestSum := array[0]

	maxSumPerIndex := make([]int, len(array))
	maxSumPerIndex[0] = array[0]

	for i := 1; i < len(array); i++ {
		maxPossibleSumAtIndex := max(array[i], maxSumPerIndex[i-1]+array[i])
		maxSumPerIndex[i] = maxPossibleSumAtIndex

		if maxPossibleSumAtIndex > greatestSum {
			greatestSum = maxPossibleSumAtIndex
		}
	}

	return greatestSum
}

func GreatestSubArraySum(array []int) int {
	greatestSum := math.MinInt

	for leftBoundary := range array {
		for rightBoundary := leftBoundary; rightBoundary < len(array); rightBoundary++ {
			sum := 0

			for i := leftBoundary; i <= rightBoundary; i++ {
				sum += array[i]
			}

			if greatestSum < sum {
				greatestSum = sum
			}
		}
	}

	return greatestSum
}
