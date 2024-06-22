package medium26

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func SymmetricalTree(tree *BinaryTree) bool {
	isSymmetrical := true

	leftInOrderNodes := []int{}
	if tree.Left != nil {
		tree.Left.traverseFromLeftToRight(&leftInOrderNodes)
	}

	rightInOrderNodes := []int{}
	if tree.Right != nil {
		tree.Right.traverseFromRightToLeft(&rightInOrderNodes)
	}

	if len(leftInOrderNodes) != len(rightInOrderNodes) {
		return false
	}

	for index, value := range leftInOrderNodes {
		if rightInOrderNodes[index] != value {
			isSymmetrical = false
			break
		}
	}

	return isSymmetrical
}

func (tree *BinaryTree) traverseFromLeftToRight(inOrderNodes *[]int) {
	if tree.Left != nil {
		tree.Left.traverseFromLeftToRight(inOrderNodes)
	}
	*inOrderNodes = append(*inOrderNodes, tree.Value)
	if tree.Right != nil {
		tree.Right.traverseFromLeftToRight(inOrderNodes)
	}
}

func (tree *BinaryTree) traverseFromRightToLeft(inOrderNodes *[]int) {
	if tree.Right != nil {
		tree.Right.traverseFromRightToLeft(inOrderNodes)
	}
	*inOrderNodes = append(*inOrderNodes, tree.Value)
	if tree.Left != nil {
		tree.Left.traverseFromRightToLeft(inOrderNodes)
	}
}
