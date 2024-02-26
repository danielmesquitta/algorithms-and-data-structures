package easy02

import (
	"sort"
)

func SortedSquaredArray(array []int) []int {
	sortedArray := make([]int, len(array))

	for _, num := range array {
		squaredNum := num * num
		sortedArray = append(sortedArray, squaredNum)
	}

	sort.Slice(sortedArray, func(i, j int) bool {
		return sortedArray[i] < sortedArray[j]
	})

	return sortedArray
}

func OptimalSortedSquaredArray(array []int) []int {
	squares := make([]int, len(array))
	start := 0
	end := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {
		if array[start]*-1 > array[end] {
			squares[i] = array[start] * array[start]
			start += 1
		} else {
			squares[i] = array[end] * array[end]
			end -= 1
		}
	}

	return squares
}
