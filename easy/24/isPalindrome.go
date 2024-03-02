package easy24

import "strings"

func IsPalindrome(str string) bool {
	result := true
	strArray := strings.Split(str, "")

	j := len(strArray) - 1
	for _, str := range strArray {
		if str != strArray[j] {
			result = false
			break
		}
		j--
	}

	return result
}
