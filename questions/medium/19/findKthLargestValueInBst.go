package medium19

type BST struct {
	Value       int
	Left, Right *BST
}

func (tree *BST) findKthLargestValueInBst(k int, array *[]int) {
	if len(*array) == k {
		return
	}

	if tree.Right != nil {
		tree.Right.findKthLargestValueInBst(k, array)
	}

	if len(*array) < k {
		*array = append(*array, tree.Value)
	}

	if tree.Left != nil {
		tree.Left.findKthLargestValueInBst(k, array)
	}
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	array := []int{}
	tree.findKthLargestValueInBst(k, &array)
	return array[len(array)-1]
}
