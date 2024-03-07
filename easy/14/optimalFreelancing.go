package easy14

import (
	"slices"
)

func setAvailableDeadline(unavailableDeadlines map[int]struct{}, deadlinePtr *int) {
	deadline := *deadlinePtr
	if _, ok := unavailableDeadlines[deadline]; ok {
		deadline--
		setAvailableDeadline(unavailableDeadlines, &deadline)
	}
	*deadlinePtr = deadline
}

func sortByBestPayments(jobs []map[string]int) {
	sortFunc := func(a, b map[string]int) int {
		return b["payment"] - a["payment"]
	}
	slices.SortFunc(jobs, sortFunc)
}

func OptimalFreelancing(jobs []map[string]int) int {
	optimalPayment := 0
	unavailable := map[int]struct{}{}
	sortByBestPayments(jobs)
	for _, job := range jobs {
		deadline := job["deadline"]
		if deadline > 7 {
			deadline = 7
		}
		setAvailableDeadline(unavailable, &deadline)
		if deadline < 1 {
			continue
		}
		unavailable[deadline] = struct{}{}
		optimalPayment += job["payment"]
	}
	return optimalPayment
}
