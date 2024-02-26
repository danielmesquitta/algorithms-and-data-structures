package easy08

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func (root *BinaryTree) countDepth(depth int) int {
	if root == nil {
		return 0
	}

	return depth + root.Left.countDepth(depth+1) + root.Right.countDepth(depth+1)
}

func NodeDepths(root *BinaryTree) int {
	result := root.countDepth(0)

	return result
}
