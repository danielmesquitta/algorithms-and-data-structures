package medium00

import (
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	answers := [][]int{}

	sort.Ints(array)

	for index, curr := range array {
		leftIndex := index + 1
		rightIndex := len(array) - 1

		if leftIndex == rightIndex {
			break
		}

		var findTriplets func()
		findTriplets = func() {
			right := array[rightIndex]
			left := array[leftIndex]

			sum := curr + left + right

			switch {
			case sum > target:
				rightIndex--

			case sum < target:
				leftIndex++

			default:
				triplets := []int{curr, left, right}
				answers = append(answers, triplets)
				leftIndex++
			}

			if leftIndex != rightIndex {
				findTriplets()
			}
		}

		findTriplets()
	}

	return answers
}
