package easy28

import "strings"

func getCountByChar(str string) map[string]int {
	charsCount := map[string]int{}
	chars := strings.Split(str, "")
	for _, char := range chars {
		charsCount[char]++
	}
	return charsCount
}

func GenerateDocument(chars string, document string) bool {
	canGenerateDocument := true

	charsCount := getCountByChar(chars)
	documentCharsCount := getCountByChar(document)

	for key, value := range documentCharsCount {
		if charsCount[key] < value {
			canGenerateDocument = false
		}
	}

	return canGenerateDocument
}
