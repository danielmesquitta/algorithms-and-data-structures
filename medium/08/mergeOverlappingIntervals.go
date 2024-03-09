package medium08

import (
	"slices"
)

func isOverlapping(target []int, comparison []int) bool {
	return target[1] >= comparison[0]
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	merged := [][]int{}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	for i := range intervals {
		if i == 0 {
			merged = append(merged, intervals[i])
		}
		if i == len(intervals)-1 {
			break
		}
		curr := merged[len(merged)-1]
		next := intervals[i+1]
		if !isOverlapping(curr, next) {
			merged = append(merged, next)
			continue
		}
		if next[1] > curr[1] {
			curr[1] = next[1]
		}
	}

	return merged
}
