package easy07

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func (root *BinaryTree) isLeaf() bool {
	return root != nil && root.Left == nil && root.Right == nil
}

func (root *BinaryTree) setBranchSums(sums *[]int, sum int) {
	if root == nil {
		return
	}

	sum += root.Value

	if root.isLeaf() {
		*sums = append(*sums, sum)
	}

	root.Left.setBranchSums(sums, sum)
	root.Right.setBranchSums(sums, sum)
}

func BranchSums(root *BinaryTree) []int {
	sums := []int{}

	root.setBranchSums(&sums, 0)

	return sums
}
