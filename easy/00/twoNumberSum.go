package easy00

/**
 * Write a function that takes in a non-empty array of distinct integers and an
 * integer representing a target sum. If any two numbers in the input array sum
 * up to the target sum, the function should return them in an array, in any
 * order. If no two numbers sum up to the target sum, the function should return
 * an empty array.
 */

// Complexity Analysis: O(n^2) time
func TwoNumberSum(array []int, target int) []int {
	for _, num1 := range array {
		for _, num2 := range array {
			if num1 == num2 {
				continue
			}

			if num1+num2 == target {
				return []int{num1, num2}
			}
		}
	}

	return []int{}
}

// Complexity Analysis: O(n) time
func BetterTwoNumberSum(array []int, target int) []int {
	storedValues := map[int]struct{}{}

	for _, currentNum := range array {
		numNeeded := target - currentNum

		if _, ok := storedValues[numNeeded]; ok {
			return []int{numNeeded, currentNum}
		}

		storedValues[currentNum] = struct{}{}
	}

	return []int{}
}
