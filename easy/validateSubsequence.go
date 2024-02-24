package easy

func IsValidSubsequence(array []int, sequence []int) bool {
	j := 0
	for i := 0; i < len(array) && j < len(sequence); i++ {
		if array[i] == sequence[j] {
			j++
		}
	}
	return j == len(sequence)
}
