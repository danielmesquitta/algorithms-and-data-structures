package medium18

type BST struct {
	Value       int
	Left, Right *BST
}

func MinHeightBST(array []int) *BST {
	tree := auxMinHeightBST(nil, array, 0, len(array))
	return tree
}

func auxMinHeightBST(tree *BST, array []int, start, end int) *BST {
	middle := getMiddle(start, end)
	node := array[middle]

	if tree == nil {
		tree = &BST{Value: node}
	} else {
		tree.Insert(node)
	}

	if middle > start {
		auxMinHeightBST(tree, array, start, middle)
	}

	if middle < end-1 {
		auxMinHeightBST(tree, array, middle+1, end)
	}

	return tree
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func getMiddle(start, end int) int {
	return (end + start) / 2
}
