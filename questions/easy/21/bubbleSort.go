package easy21

func BubbleSort(array []int) []int {
	left, right, end := 0, 1, len(array)-1

	swappedDuringCycle := false

	for end != 0 {
		if array[left] > array[right] {
			array[left], array[right] = array[right], array[left]
			swappedDuringCycle = true
		}

		if finishedCycle := right == end; finishedCycle {
			if !swappedDuringCycle {
				break
			}

			end--
			left = 0
			right = 1
			continue
		}

		left++
		right++
	}

	return array
}
