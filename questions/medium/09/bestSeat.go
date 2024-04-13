package medium09

import "math"

type Bounds struct {
	Start int
	End   int
}

func (b *Bounds) IsEmpty() bool {
	return b.IsEmptyStart() && b.IsEmptyEnd()
}

func (b *Bounds) IsEmptyStart() bool {
	return b.Start == -1
}

func (b *Bounds) IsEmptyEnd() bool {
	return b.End == -1
}

func (b *Bounds) GetMiddle() int {
	return (b.Start + (b.End - 1)) / 2
}

func NewBounds() *Bounds {
	return &Bounds{
		Start: -1,
		End:   -1,
	}
}

func GetLargestZeroSubarrayBounds(array []int) *Bounds {
	bounds := []*Bounds{}

	for i, num := range array {
		if len(bounds) == 0 {
			bounds = append(bounds, NewBounds())
		}

		last := bounds[len(bounds)-1]
		switch num {
		case 1:
			if last.IsEmpty() {
				continue
			}
			bounds = append(bounds, NewBounds())
		case 0:
			if last.IsEmptyStart() {
				last.Start = i
			}
			last.End = i + 1
		}
	}

	largest := math.MinInt
	largestIndex := -1
	for i, bound := range bounds {
		length := bound.End - bound.Start
		if length > largest {
			largest = length
			largestIndex = i
		}
	}

	if largestIndex == -1 {
		return NewBounds()
	}

	bound := bounds[largestIndex]

	return bound
}

func BestSeat(seats []int) int {
	bound := GetLargestZeroSubarrayBounds(seats)

	if bound.IsEmpty() {
		return -1
	}

	return bound.GetMiddle()
}
