package easy12

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	blueIsTaller := true
	redIsTaller := true
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	for i, red := range redShirtHeights {
		blue := blueShirtHeights[i]
		if blue >= red {
			redIsTaller = false
		}
		if red >= blue {
			blueIsTaller = false
		}
		if !redIsTaller && !blueIsTaller {
			break
		}
	}
	return blueIsTaller || redIsTaller
}
