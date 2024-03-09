package medium07

func FirstDuplicateValue(array []int) int {
	visited := map[int]struct{}{}
	for _, num := range array {
		if _, ok := visited[num]; ok {
			return num
		}
		visited[num] = struct{}{}
	}
	return -1
}
