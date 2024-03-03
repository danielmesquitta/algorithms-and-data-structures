package easy30

import (
	"strings"
)

func reverse(str string) string {
	chars := strings.Split(str, "")
	reversedChars := make([]string, len(chars))
	for i := len(chars) - 1; i >= 0; i-- {
		reversedChars = append(reversedChars, chars[i])
	}
	return strings.Join(reversedChars, "")
}

func Semordnilap(words []string) [][]string {
	pairsMap := map[string]string{}
	pairs := [][]string{}
	for _, word := range words {
		if reversed, ok := pairsMap[word]; ok {
			pairs = append(pairs, []string{word, reversed})
			continue
		}
		pairsMap[reverse(word)] = word
	}

	return pairs
}
