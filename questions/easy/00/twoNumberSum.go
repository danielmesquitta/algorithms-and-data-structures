package easy00

// Complexity Analysis: O(n^2) time
func BruteForceTwoNumberSum(array []int, target int) []int {
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
func TwoNumberSum(array []int, target int) []int {
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
