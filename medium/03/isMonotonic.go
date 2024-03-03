package medium03

func IsMonotonic(array []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i, curr := range array {
		if i == len(array)-1 {
			break
		}
		next := array[i+1]

		if curr > next {
			isIncreasing = false
		}

		if curr < next {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
