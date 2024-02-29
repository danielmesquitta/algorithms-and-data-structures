package easy14

import (
	"slices"
)

func OptimalFreelancing(jobs []map[string]int) int {
	totalPayment := 0

	// Sort from greatest to lowest payment
	slices.SortFunc(jobs, func(a, b map[string]int) int {
		return b["payment"] - a["payment"]
	})

	occupiedDeadlines := map[int]struct{}{}

	for _, job := range jobs {
		deadline := job["deadline"]

		if deadline > 7 {
			deadline = 7
		}

		var findDeadlineAvailable func()
		findDeadlineAvailable = func() {
			if _, ok := occupiedDeadlines[deadline]; ok {
				deadline--
				findDeadlineAvailable()
			}
		}
		findDeadlineAvailable()

		if deadline < 1 {
			continue
		}

		occupiedDeadlines[deadline] = struct{}{}
		totalPayment += job["payment"]
	}

	return totalPayment
}
