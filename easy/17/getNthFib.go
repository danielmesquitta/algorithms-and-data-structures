package easy17

func GetNthFib(n int) int {
	left, right := 0, 1

	if n == 1 {
		return left
	}
	if n == 2 {
		return right
	}

	num := 0
	for i := 0; i < n-2; i++ {
		num = left + right
		left = right
		right = num
	}

	return num
}
