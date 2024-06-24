package medium27

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func SplitBinaryTree(tree *BinaryTree) int {
	foundRightSums, foundLeftSums := map[int]struct{}{}, map[int]struct{}{}
	var rightSubtreeSum, leftSubtreeSum int

	if tree.Left != nil {
		leftSubtreeSum = tree.Left.registerSubtreeSums(&foundLeftSums)
	}

	if tree.Right != nil {
		rightSubtreeSum = tree.Right.registerSubtreeSums(&foundRightSums)
	}

	treeSum := tree.Value + leftSubtreeSum + rightSubtreeSum

	if isOdd := treeSum%2 == 1; isOdd {
		return 0
	}

	desiredSum := treeSum / 2

	_, foundDesiredSumOnLeft := foundLeftSums[desiredSum]
	_, foundDesiredSumOnRight := foundRightSums[desiredSum]

	if foundDesiredSumOnLeft || foundDesiredSumOnRight {
		return desiredSum
	}

	return 0
}

func (tree *BinaryTree) registerSubtreeSums(
	foundSums *map[int]struct{},
) (subtreeSum int) {
	if tree == nil {
		return 0
	}

	leftSum := tree.Left.registerSubtreeSums(foundSums)
	rightSum := tree.Right.registerSubtreeSums(foundSums)

	subtreeSum = tree.Value + leftSum + rightSum
	(*foundSums)[subtreeSum] = struct{}{}

	return subtreeSum
}
