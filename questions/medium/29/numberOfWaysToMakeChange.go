package medium29

func NumberOfWaysToMakeChange(target int, denoms []int) int {
	ways := make([]int, target+1)
	ways[0] = 1

	for _, denom := range denoms {
		for i := denom; i <= target; i++ {
			if denom > i {
				continue
			}

			ways[i] += ways[i-denom]
		}
	}

	return ways[len(ways)-1]
}
