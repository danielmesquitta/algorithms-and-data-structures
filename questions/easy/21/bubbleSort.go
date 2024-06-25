package easy21

func BubbleSort(array []int) []int {
	end := len(array) - 1
	for end != 0 {
		swapped := false
		for left := 0; left < end; left++ {
			right := left + 1
			if array[left] > array[right] {
				array[left], array[right] = array[right], array[left]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		end--
	}
	return array
}
