package medium10

func getSubArraysFromArray(array []int) [][]int {
	subArrays := [][]int{}

	for i := range array {
		for j := range array {
			if i > j {
				continue
			}
			subArrays = append(subArrays, array[i:j+1])
		}
	}

	return subArrays
}

func sumArray(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

func BruteForceZeroSumSubarray(nums []int) bool {
	subArrays := getSubArraysFromArray(nums)

	for _, array := range subArrays {
		sum := sumArray(array)
		if sum == 0 {
			return true
		}
	}

	return false
}

func ZeroSumSubarray(nums []int) bool {
	sums := map[int]struct{}{0: {}}
	currentSum := 0

	for _, num := range nums {
		currentSum += num
		if _, ok := sums[currentSum]; ok {
			return true
		}
		sums[currentSum] = struct{}{}
	}

	return false
}
