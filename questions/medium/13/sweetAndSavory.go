package medium13

import (
	"math"
	"sort"
)

func SweetAndSavory(dishes []int, target int) []int {
	sort.Ints(dishes)

	pair := make([]int, 2)
	closestDistance := math.MaxInt

	left, right := 0, len(dishes)-1

loop:
	for left < right && dishes[left] < 0 && dishes[right] > 0 {
		pairFlavor := dishes[left] + dishes[right]
		switch {
		case pairFlavor > target:
			right--

		case pairFlavor < target:
			distance := AbsInt(pairFlavor - target)

			if closestDistance > distance {
				closestDistance = distance
				pair = []int{dishes[left], dishes[right]}
			}

			left++

		default:
			pair = []int{dishes[left], dishes[right]}
			break loop
		}
	}

	return pair
}

func AbsInt(num int) int {
	if num < 0 {
		num = num * -1
	}
	return num
}
