package medium12

func MajorityElement(array []int) int {
	occurrences := map[int]int{}

	majorityElementOccurrences := -1
	for _, num := range array {
		occurrences[num]++

		if occurrences[num] > majorityElementOccurrences {
			majorityElementOccurrences = occurrences[num]
		}
	}

	majorityElement := -1
	for element, count := range occurrences {
		if count == majorityElementOccurrences {
			majorityElement = element
			break
		}
	}

	return majorityElement
}
