package easy23

import "math"

func SelectionSort(array []int) []int {
	i, start, end := 0, 0, len(array)-1
	selectedI, selectedVal := -1, math.MaxInt
	for start <= end {
		if array[i] < selectedVal {
			selectedVal = array[i]
			selectedI = i
		}
		if i == end {
			array[start], array[selectedI] = array[selectedI], array[start]
			start++
			i = start
			selectedI, selectedVal = -1, math.MaxInt
			continue
		}
		i++
	}
	return array
}
