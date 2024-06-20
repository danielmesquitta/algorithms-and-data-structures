package medium20

import (
	"math"
)

type BST struct {
	Value       int
	Left, Right *BST
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	max := math.MaxInt
	tree, _ := makeNode(preOrderTraversalValues, 0, max)
	return tree
}

func makeNode(array []int, index, max int) (*BST, int) {
	if index >= len(array) {
		return nil, index
	}

	if value := array[index]; value < max {
		node := &BST{Value: value}
		index++

		node.Left, index = makeNode(array, index, value)
		node.Right, index = makeNode(array, index, max)
		return node, index
	}

	return nil, index
}
