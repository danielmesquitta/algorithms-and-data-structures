package medium04

type Direction byte

const (
	Right Direction = iota
	Down
	Left
	Up
)

func SpiralTraverse(matrix [][]int) []int {
	array := []int{}

	rowsCount := len(matrix)
	colsCount := len(matrix[0])
	size := rowsCount * colsCount

	startRow, endRow := 0, rowsCount-1
	startCol, endCol := 0, colsCount-1

	var direction Direction
	switch colsCount {
	case 1:
		direction = Down
	default:
		direction = Right
	}

	row, col := 0, 0
	for len(array) != size {
		array = append(array, matrix[row][col])

		switch direction {
		case Right:
			if col < endCol {
				col++
			}
			if col == endCol {
				startRow++
				direction = Down
			}
			continue

		case Down:
			if row < endRow {
				row++
			}
			if row == endRow {
				endCol--
				direction = Left
			}
			continue

		case Left:
			if col > startCol {
				col--
			}
			if col == startCol {
				endRow--
				direction = Up
			}
			continue

		case Up:
			if row > startRow {
				row--
			}
			if row == startRow {
				startCol++
				direction = Right
			}
			continue
		}
	}

	return array
}
