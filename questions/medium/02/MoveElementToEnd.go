package medium02

func MoveElementToEnd(array []int, toMove int) []int {
	end := len(array) - 1

	for i := range array {
		if i >= end {
			break
		}
		for array[i] == toMove && i < end {
			array[end], array[i] = array[i], array[end]
			end--
		}
	}

	return array
}
