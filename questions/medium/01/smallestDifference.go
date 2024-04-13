package medium01

import (
	"math"
	"sort"
)

func BruteForceSmallestDifference(array1, array2 []int) [2]int {
	storedDistances := map[int][2]int{}
	smallestDifference := math.MaxInt32

	for _, num1 := range array1 {
		for _, num2 := range array2 {
			distance := int(math.Abs(float64(num1 - num2)))
			if distance < smallestDifference {
				smallestDifference = distance
			}
			storedDistances[distance] = [2]int{num1, num2}
		}
	}

	return storedDistances[smallestDifference]
}

func SmallestDifference(array1, array2 []int) [2]int {
	storedDifferences := map[int][2]int{}
	smallestDifference := math.MaxInt32

	sort.Ints(array1)
	sort.Ints(array2)

	var findSmallestDistance func(i, j int)
	findSmallestDistance = func(i, j int) {
		if i == len(array1) || j == len(array2) {
			return
		}

		num1 := array1[i]
		num2 := array2[j]

		distance := num1 - num2

		absDistance := int(math.Abs(float64(distance)))
		storedDifferences[absDistance] = [2]int{num1, num2}

		if absDistance < smallestDifference {
			smallestDifference = absDistance
		}

		if absDistance == 0 {
			return
		}

		if num1 > num2 {
			findSmallestDistance(i, j+1)
		} else {
			findSmallestDistance(i+1, j)
		}
	}

	findSmallestDistance(0, 0)

	return storedDifferences[smallestDifference]
}
