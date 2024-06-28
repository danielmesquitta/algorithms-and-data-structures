package medium32

func NumberOfWaysToTraverseGraph(width int, height int) int {
	matrix := make([][]int, height)

	for i := range matrix {
		matrix[i] = make([]int, width)

		for j := range matrix[i] {
			if i == 0 || j == 0 {
				matrix[i][j] = 1
				continue
			}

			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	return matrix[height-1][width-1]
}
