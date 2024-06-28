package medium30

const errValue = -1

func MinNumberOfCoinsForChange(target int, coins []int) int {
	minNumCoinsPerChange := make([]int, target+1)
	minNumCoinsPerChange[0] = 0

	for i := 1; i <= target; i++ {
		minNumCoinsPerChange[i] = errValue
	}

	for change := 1; change <= target; change++ {
		for _, coin := range coins {
			if coin > change {
				continue
			}

			proposedSubtotal := minNumCoinsPerChange[change-coin]
			if proposedSubtotal == errValue {
				continue
			}

			currentBest := minNumCoinsPerChange[change]
			proposedBest := proposedSubtotal + 1

			if currentBest == errValue || (proposedBest <= currentBest) {
				minNumCoinsPerChange[change] = proposedBest
			}
		}
	}

	return minNumCoinsPerChange[target]
}
