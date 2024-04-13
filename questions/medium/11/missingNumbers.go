package medium11

import "sort"

func appendInIndex(array []int, value int, index int) []int {
	left, right := index, index+1

	tail := nums[right : end+1]
}

func MissingNumbers(nums []int) []int {
	sort.Ints(nums)
	missing := []int{}

	left, right, end := 0, 1, len(nums)-1
	for len(missing) < 2 {
		curr, next := nums[left], nums[right]
		if curr+1 != next {
			tail := append([]int{curr + 1}, nums[right:end+1]...)
			nums = append(nums[0:left+1], tail...)
		}
	}

	return missing
}
