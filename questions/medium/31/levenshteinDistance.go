package medium31

func LevenshteinDistance(a, b string) int {
	la := len(a)
	lb := len(b)

	distances := make([][]int, la+1)
	for i := range distances {
		distances[i] = make([]int, lb+1)
	}

	for i := 0; i <= la; i++ {
		distances[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		distances[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			distances[i][j] = min(
				distances[i-1][j]+1,      // Deletion
				distances[i][j-1]+1,      // Insertion
				distances[i-1][j-1]+cost, // Substitution
			)
		}
	}

	return distances[la][lb]
}
