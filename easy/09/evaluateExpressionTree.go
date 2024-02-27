package easy09

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

const (
	ADDITION       = -1
	SUBTRACTION    = -2
	DIVISION       = -3
	MULTIPLICATION = -4
)

func (root *BinaryTree) isLeaf() bool {
	return root != nil && root.Left == nil && root.Right == nil
}

func (root *BinaryTree) evaluateExpression() int {
	if root.isLeaf() {
		return root.Value
	}

	switch root.Value {
	case ADDITION:
		return root.Left.evaluateExpression() + root.Right.evaluateExpression()
	case SUBTRACTION:
		return root.Left.evaluateExpression() - root.Right.evaluateExpression()
	case DIVISION:
		return root.Left.evaluateExpression() / root.Right.evaluateExpression()
	case MULTIPLICATION:
		return root.Left.evaluateExpression() * root.Right.evaluateExpression()
	default:
		return 0
	}
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	result := tree.evaluateExpression()

	return result
}
