package medium24

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	var leftDepth, rightDepth int
	isBalanced, _ := tree.auxHeightBalancedBinaryTree(&leftDepth, &rightDepth)

	return isBalanced
}

func (tree *BinaryTree) auxHeightBalancedBinaryTree(
	leftDepth, rightDepth *int,
) (bool, int) {
	if tree.Left != nil {
		*leftDepth++

		tree.Left.auxHeightBalancedBinaryTree(
			leftDepth,
			rightDepth,
		)
	}

	if tree.Right != nil {
		*rightDepth++

		tree.Right.auxHeightBalancedBinaryTree(
			leftDepth,
			rightDepth,
		)
	}

	difference := Abs(*leftDepth - *rightDepth)

	return difference <= 1, 1
}

func Abs(n int) int {
	if n < 0 {
		return (n * -1)
	}
	return n
}
