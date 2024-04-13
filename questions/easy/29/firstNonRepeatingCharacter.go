package easy29

import (
	"strings"
)

func FirstNonRepeatingCharacter(str string) int {
	charactersCount := map[string]int{}

	chars := strings.Split(str, "")

	for _, char := range chars {
		charactersCount[char]++
	}

	for i, char := range chars {
		if charactersCount[char] == 1 {
			return i
		}
	}

	return -1
}
