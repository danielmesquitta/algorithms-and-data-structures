package easy17

func GetNthFib(n int) int {
	left, right := 0, 1
	if n == 1 {
		return left
	}
	for i := 0; i < n-2; i++ {
		left, right = right, left+right
	}
	return right
}
