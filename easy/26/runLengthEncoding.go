package easy26

import (
	"fmt"
	"strings"
)

func encode(encodedPtr *string, countPtr *int, char string) {
	*encodedPtr = fmt.Sprintf("%s%d%s", *encodedPtr, *countPtr, char)
	*countPtr = 1
}

func RunLengthEncoding(str string) string {
	encoded := ""
	charArray := strings.Split(str, "")

	count := 1
	for i, curr := range charArray {
		if len(charArray) == i+1 || count == 9 {
			encode(&encoded, &count, curr)
			continue
		}

		next := charArray[i+1]
		if curr == next {
			count++
			continue
		}

		encode(&encoded, &count, curr)
	}

	return encoded
}
