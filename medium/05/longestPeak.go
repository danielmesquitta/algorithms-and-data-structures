// package medium05

// import (
// 	"math"
// )

// func FindPeaks(array []int) []int {
// 	peakIndexes := []int{}

// 	for i := 1; i < len(array)-1; i++ {
// 		prev := array[i-1]
// 		curr := array[i]
// 		next := array[i+1]

// 		foundPeak := prev < curr && curr > next
// 		if foundPeak {
// 			peakIndexes = append(peakIndexes, i)
// 		}
// 	}

// 	return peakIndexes
// }

// func countLeftStrictlyLower(array []int, index int) int {
// }

// func LongestPeak(array []int) int {
// 	peakIndexes := FindPeaks(array)

// 	longestPeak := math.MinInt
// 	for _, index := range peakIndexes {
// 		prev := array[index-1]
// 		next := array[index+1]
// 	}

// 	if longestPeak == math.MinInt {
// 		return 0
// 	}

// 	return longestPeak
// }
