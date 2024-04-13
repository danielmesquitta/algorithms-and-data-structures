package medium06

func ArrayOfProductsBruteForce(array []int) []int {
	output := make([]int, len(array))
	for i := range array {
		output[i] = 1
		for j, next := range array {
			if i == j {
				continue
			}
			output[i] *= next
		}
	}
	return output
}

func multiplyByOthers(output *[]int, input []int, oIndex int, iIndex int) {
	if oIndex == -1 || oIndex == len(input) {
		return
	}
	if oIndex != iIndex {
		(*output)[oIndex] *= input[iIndex]
	}
	if oIndex <= iIndex {
		multiplyByOthers(output, input, oIndex-1, iIndex)
	}
	if oIndex >= iIndex {
		multiplyByOthers(output, input, oIndex+1, iIndex)
	}
}

func ArrayOfProducts(input []int) []int {
	output := make([]int, len(input))
	for i := range output {
		output[i] = 1
	}
	for i := range input {
		multiplyByOthers(&output, input, i, i)
	}

	return output
}
