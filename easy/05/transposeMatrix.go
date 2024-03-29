package easy05

func TransposeMatrix(matrix [][]int) [][]int {
	transposedMatrix := make([][]int, len(matrix[0]))
	for col := 0; col < len(matrix[0]); col++ {
		transposedMatrix[col] = make([]int, len(matrix))
		for row := 0; row < len(matrix); row++ {
			transposedMatrix[col][row] = matrix[row][col]
		}
	}
	return transposedMatrix
}
