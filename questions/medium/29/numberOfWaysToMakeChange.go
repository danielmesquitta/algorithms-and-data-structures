package medium29

func NumberOfWaysToMakeChange(target int, coins []int) int {
	waysPerAmount := make([]int, target+1)
	waysPerAmount[0] = 1

	for _, coin := range coins {
		for amount := coin; amount <= target; amount++ {
			waysPerAmount[amount] += waysPerAmount[amount-coin]
		}
	}

	return waysPerAmount[len(waysPerAmount)-1]
}
