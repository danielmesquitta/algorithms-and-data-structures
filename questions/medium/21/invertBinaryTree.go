package medium21

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	tree.Left, tree.Right = tree.Right, tree.Left

	if tree.Left != nil {
		tree.Left.InvertBinaryTree()
	}

	if tree.Right != nil {
		tree.Right.InvertBinaryTree()
	}
}
