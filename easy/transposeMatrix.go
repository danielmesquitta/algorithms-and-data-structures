package easy

func TransposeMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}

	transposedMatrix := make([][]int, len(matrix[0]))

	for col := 0; col < len(matrix[0]); col++ {
		transposedMatrix[col] = make([]int, len(matrix))

		for row := 0; row < len(matrix); row++ {
			transposedMatrix[col][row] = matrix[row][col]
		}
	}

	return transposedMatrix
}
