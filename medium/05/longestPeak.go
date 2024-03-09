package medium05

import (
	"math"
)

type Direction byte

// =================
// Brute force
// =================

const (
	GoingUp Direction = iota
	GoingDown
)

func LongestPeakBruteForce(array []int) int {
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

// =================
// Better solution
// =================

const (
	Right Direction = iota
	Left
)

func findPeaks(array []int) []int {
	peakIndexes := []int{}

	for i := 1; i < len(array)-1; i++ {
		prev := array[i-1]
		curr := array[i]
		next := array[i+1]

		foundPeak := prev < curr && curr > next
		if foundPeak {
			peakIndexes = append(peakIndexes, i)
		}
	}

	return peakIndexes
}

func countStrictlyLower(array []int, index int, direction Direction) int {
	count := 0
	isBetweenSecondAndPenultimate := func(i int) bool {
		return i > 0 && i < len(array)-1
	}

	step := -1
	if direction == Right {
		step = 1
	}

	for isBetweenSecondAndPenultimate(index) {
		curr := array[index]
		next := array[index+step]
		if next >= curr {
			break
		}
		count++
		index += step
	}

	return count
}

func LongestPeak(array []int) int {
	peakIndexes := findPeaks(array)

	longestPeak := math.MinInt
	for _, index := range peakIndexes {
		left := countStrictlyLower(array, index, Left)
		right := countStrictlyLower(array, index, Right)

		peakLength := left + right + 1
		if peakLength > longestPeak {
			longestPeak = peakLength
		}
	}

	if longestPeak == math.MinInt {
		return 0
	}

	return longestPeak
}
