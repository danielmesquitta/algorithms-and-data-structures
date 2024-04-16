package medium11

func MissingNumbers(nums []int) []int {
	missing := []int{}
	missingMaxLength := 2
	n := len(nums) + missingMaxLength

	foundNums := map[int]struct{}{}
	for _, num := range nums {
		foundNums[num] = struct{}{}
	}

	for i := range n {
		num := i + 1
		if len(missing) == missingMaxLength {
			continue
		}

		if _, ok := foundNums[num]; !ok {
			missing = append(missing, num)
		}
	}

	return missing
}
