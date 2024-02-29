package easy13

import (
	"slices"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	slices.Sort(redShirtSpeeds)

	if fastest {
		slices.SortFunc(blueShirtSpeeds, func(a, b int) int {
			return b - a
		})
	} else {
		slices.Sort(blueShirtSpeeds)
	}

	sum := 0
	for i, red := range redShirtSpeeds {
		blue := blueShirtSpeeds[i]
		if red > blue {
			sum += red
		} else {
			sum += blue
		}
	}

	return sum
}
