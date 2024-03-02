package easy27

import "strings"

func CommonCharacters(strs []string) []string {
	foundChars := map[string]int{}

	for i, str := range strs {
		chars := strings.Split(str, "")
		for _, char := range chars {
			if foundChars[char] == i {
				foundChars[char]++
			}
		}
	}

	response := []string{}
	length := len(strs)
	for key, value := range foundChars {
		if value == length {
			response = append(response, key)
		}
	}

	return response
}
