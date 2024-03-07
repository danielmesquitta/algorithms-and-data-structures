package easy13

import (
	"slices"
)

func pairByShirts(redSpeeds []int, blueSpeeds []int, fastest bool) {
	slices.Sort(redSpeeds)
	if fastest {
		slices.SortFunc(blueSpeeds, func(a, b int) int {
			return b - a
		})
	} else {
		slices.Sort(blueSpeeds)
	}
}

func TandemBicycle(redSpeeds []int, blueSpeeds []int, fastest bool) int {
	pairByShirts(redSpeeds, blueSpeeds, fastest)
	sum := 0
	for i, red := range redSpeeds {
		blue := blueSpeeds[i]
		if red > blue {
			sum += red
		} else {
			sum += blue
		}
	}
	return sum
}
