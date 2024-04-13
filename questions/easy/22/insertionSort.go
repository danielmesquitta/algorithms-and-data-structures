package easy22

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		current := array[i]
		j := i - 1
		for j >= 0 && array[j] > current {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = current
	}
	return array
}
