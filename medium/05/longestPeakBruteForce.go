package medium05

import "math"

type Direction byte

const (
	GoingUp Direction = iota
	GoingDown
)

func LongestPeak(array []int) int {
	peaks := []int{}
	index := 0
	direction := GoingUp

	for i, curr := range array {
		isInPeak := len(peaks) > index

		if i == len(array)-1 {
			if direction == GoingUp && isInPeak {
				// it is not a peak, delete it
				peaks = peaks[0:index]
			}
			break
		}

		next := array[i+1]

		switch direction {
		case GoingUp:
			switch {
			case curr < next:
				// add 2 if starting peak,
				// add 1 if continuing peak
				if !isInPeak {
					peaks = append(peaks, 1)
				}
				peaks[index]++
			case curr > next:
				// change direction and add 1
				if isInPeak {
					peaks[index]++
					direction = GoingDown
				}
			case curr == next:
				// it is not a peak, delete it
				if isInPeak {
					peaks = peaks[0:index]
				}
			}

		case GoingDown:
			switch {
			case curr > next:
				// add
				peaks[index]++
			case curr == next:
				// finish
				direction = GoingUp
				index++
			case curr < next:
				//finish and start new peak
				direction = GoingUp
				index++
				if index == len(peaks) {
					peaks = append(peaks, 1)
				}
				peaks[index]++
			}
		}
	}

	longestPeak := math.MinInt
	for _, peak := range peaks {
		if peak > longestPeak {
			longestPeak = peak
		}
	}

	if longestPeak == math.MinInt {
		return 0
	}

	return longestPeak
}
