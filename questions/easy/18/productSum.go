package easy18

type SpecialArray []interface{}

func getProductSum(depth int, array SpecialArray) int {
	result := 0
	for _, item := range array {
		num, ok := item.(int)
		deeperResult := 0
		if !ok {
			deeperResult = getProductSum(depth+1, item.(SpecialArray))
		}
		result += (num + ((depth + 1) * deeperResult))
	}
	return result
}

func ProductSum(array SpecialArray) int {
	depth := 1
	return getProductSum(depth, array)
}
