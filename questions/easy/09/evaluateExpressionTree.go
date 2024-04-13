package easy09

type Node struct {
	Value       int
	Left, Right *Node
}

const (
	ADDITION       = -1
	SUBTRACTION    = -2
	DIVISION       = -3
	MULTIPLICATION = -4
)

func (n *Node) isLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) Evaluate() int {
	if n.isLeaf() {
		return n.Value
	}
	switch n.Value {
	case ADDITION:
		return n.Left.Evaluate() + n.Right.Evaluate()
	case SUBTRACTION:
		return n.Left.Evaluate() - n.Right.Evaluate()
	case DIVISION:
		return n.Left.Evaluate() / n.Right.Evaluate()
	default:
		return n.Left.Evaluate() * n.Right.Evaluate()
	}
}
