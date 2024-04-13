package easy06

import "math"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	closest := tree.Value
	for tree != nil {
		closestDistance := math.Abs(float64(closest - target))
		treeDistance := math.Abs(float64(tree.Value - target))
		if closestDistance > treeDistance {
			closest = tree.Value
		}
		if target < tree.Value {
			tree = tree.Left
		} else if target > tree.Value {
			tree = tree.Right
		} else {
			break
		}
	}
	return closest
}
