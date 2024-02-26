package easy04

import (
	"sort"
)

func generateAllSums(nums []int) []int {
	sumSet := map[int]struct{}{}

	var subsets func(index int, currentSum int)
	subsets = func(index int, currentSum int) {
		if index == len(nums) {
			sumSet[currentSum] = struct{}{}
			return
		}
		subsets(index+1, currentSum+nums[index])
		subsets(index+1, currentSum)
	}

	subsets(0, 0)

	var sums []int
	for sum := range sumSet {
		sums = append(sums, sum)
	}

	sort.Ints(sums)

	return sums
}

func NonConstructibleChange(coins []int) int {
	allSums := generateAllSums(coins)

	for index, curr := range allSums {
		if isLastItem := index+1 == len(allSums); isLastItem {
			return curr + 1
		} else {
			next := allSums[index+1]
			differenceFromNext := next - curr

			if differenceFromNext > 1 {
				return curr + 1
			}
		}
	}

	return -1
}

func ImprovedNonConstructibleChange(coins []int) int {
	sort.Slice(coins, func(a, b int) bool {
		return coins[a] < coins[b]
	})

	var sum int
	for _, elem := range coins {
		if elem > sum+1 {
			break
		}

		sum += elem
	}

	return sum + 1
}
