package medium22

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func BinaryTreeDiameter(root *BinaryTree) int {
	greaterPath := 0
	binaryTreeDiameter(root, &greaterPath)
	return greaterPath
}

func binaryTreeDiameter(node *BinaryTree, greaterPath *int) int {
	if node == nil {
		return 0
	}

	leftBranch := binaryTreeDiameter(node.Left, greaterPath)
	rightBranch := binaryTreeDiameter(node.Right, greaterPath)

	*greaterPath = Max(leftBranch+rightBranch, *greaterPath)

	return Max(leftBranch, rightBranch) + 1
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
