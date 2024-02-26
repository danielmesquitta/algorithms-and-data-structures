package easy06

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	if tree == nil {
		return 0
	}

	closestValue := tree.Value

	for tree != nil {
		closestValueDistance := Abs(closestValue - target)
		treeValueDistance := Abs(tree.Value - target)

		if closestValueDistance > treeValueDistance {
			closestValue = tree.Value
		}

		if target < tree.Value {
			tree = tree.Left
		} else if target > tree.Value {
			tree = tree.Right
		} else {
			break
		}
	}

	return closestValue
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
