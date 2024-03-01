package easy19

func BinarySearch(array []int, target int) int {
	start, end := 0, len(array)-1

	for start <= end {
		mid := (start + end) / 2
		item := array[mid]

		switch {
		case item > target:
			end = mid - 1
		case item < target:
			start = mid + 1
		default:
			return mid
		}
	}

	return -1
}
