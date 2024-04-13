package easy07

type Node struct {
	Value       int
	Left, Right *Node
}

func (n *Node) IsLeaf() bool {
	return n != nil && n.Left == nil && n.Right == nil
}

func (n *Node) setBranchSums(sums *[]int, sum int) {
	if n == nil {
		return
	}
	sum += n.Value
	if n.IsLeaf() {
		*sums = append(*sums, sum)
	}
	n.Left.setBranchSums(sums, sum)
	n.Right.setBranchSums(sums, sum)
}

func (n *Node) GetBranchSums() []int {
	sums := []int{}
	n.setBranchSums(&sums, 0)
	return sums
}
