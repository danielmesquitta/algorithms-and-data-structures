package easy20

import "math"

func FindThreeLargestNumbers(array []int) []int {
	n1st, n2nd, n3rd := math.MinInt, math.MinInt, math.MinInt

	for _, num := range array {
		if num > n1st {
			n3rd, n2nd, n1st = n2nd, n1st, num
			continue
		} else if num > n2nd {
			n3rd, n2nd = n2nd, num
			continue
		} else if num > n3rd {
			n3rd = num
		}
	}

	return []int{n3rd, n2nd, n1st}
}
