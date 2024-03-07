package easy08

type Node struct {
	Value       int
	Left, Right *Node
}

func (n *Node) countDepth(depth int) int {
	if n == nil {
		return 0
	}
	return depth + n.Left.countDepth(depth+1) + n.Right.countDepth(depth+1)
}

func (n *Node) GetNodeDepths() int {
	result := n.countDepth(0)
	return result
}
