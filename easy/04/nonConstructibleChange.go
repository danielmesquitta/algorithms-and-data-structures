package easy04

import (
	"math"
	"slices"
	"sort"
)

func generateCombinations(arr []int) [][]int {
	combinations := [][]int{}
	n := len(arr)
	numOfCombinations := int(math.Pow(2, float64(n)))
	for i := 1; i < numOfCombinations; i++ {
		combination := []int{}
		currentCombination := i
		for j := 0; j < n; j++ {
			if currentCombination%2 == 1 {
				combination = append(combination, arr[j])
			}
			currentCombination /= 2
		}
		combinations = append(combinations, combination)
	}
	return combinations
}

func sumCombinations(combinations [][]int) []int {
	uniqueSumsMap := map[int]struct{}{}
	for _, combination := range combinations {
		sum := 0
		for _, num := range combination {
			sum += num
		}
		uniqueSumsMap[sum] = struct{}{}
	}
	uniqueSums := make([]int, len(uniqueSumsMap))
	for sum := range uniqueSumsMap {
		uniqueSums = append(uniqueSums, sum)
	}
	return uniqueSums
}

func BruteForceNonConstructibleChange(coins []int) int {
	combinations := generateCombinations(coins)
	sums := sumCombinations(combinations)
	slices.Sort(sums)
	for index, curr := range sums {
		if index+1 == len(sums) {
			return curr + 1
		}
		next := sums[index+1]
		difference := next - curr
		if difference > 1 {
			return curr + 1
		}
	}
	return -1
}

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	sum := 0
	for _, coin := range coins {
		if coin > sum+1 {
			break
		}
		sum += coin
	}
	return sum + 1
}
