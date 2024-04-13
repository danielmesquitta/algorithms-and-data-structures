package easy25

func CaesarCipherEncryptor(str string, key int) string {
	strRunes := []rune(str)
	for i, char := range strRunes {
		strRunes[i] = rune((int(char)-97+key)%26 + 97)
	}
	return string(strRunes)
}
