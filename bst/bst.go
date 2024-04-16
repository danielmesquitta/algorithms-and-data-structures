package bst

import "fmt"

type BST struct {
	Value       int
	Left, Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if tree == nil {
		return &BST{Value: value}
	}
	if value < tree.Value {
		tree.Left = tree.Left.Insert(value)
	} else {
		tree.Right = tree.Right.Insert(value)
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	switch {
	case tree == nil:
		return false

	case value > tree.Value:
		return tree.Right.Contains(value)

	case value < tree.Value:
		return tree.Left.Contains(value)

	default:
		return true
	}
}

func (tree *BST) Remove(value int) *BST {
	if tree == nil {
		return tree
	}

	if value < tree.Value {
		tree.Left = tree.Left.Remove(value)
	} else if value > tree.Value {
		tree.Right = tree.Right.Remove(value)
	} else {
		if tree.Left == nil && tree.Right == nil {
			return nil
		} else if tree.Left != nil && tree.Right != nil {
			smallest := tree.Right.Smallest()
			tree.Value = smallest
			tree.Right = tree.Right.Remove(smallest)
		} else if tree.Right == nil {
			tree.Value = tree.Left.Value
			tree.Right = tree.Left.Right
			tree.Left = tree.Left.Left
		} else {
			tree.Value = tree.Right.Value
			tree.Left = tree.Right.Left
			tree.Right = tree.Right.Right
		}
	}

	return tree
}

func (tree *BST) Smallest() int {
	if tree == nil {
		panic("nil tree")
	}
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.Smallest()
}

func (tree *BST) Validate() bool {
	isValid := true

	var validate func(*BST, *bool)
	validate = func(node *BST, isValidPtr *bool) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Value >= node.Value {
			*isValidPtr = false
			return
		}
		if node.Right != nil && node.Right.Value < node.Value {
			*isValidPtr = false
			return
		}
		validate(node.Left, isValidPtr)
		validate(node.Right, isValidPtr)
	}

	validate(tree, &isValid)

	return isValid
}

func (tree *BST) Print() {
	count := 5

	var printBST func(*BST, int)
	printBST = func(node *BST, space int) {
		if node == nil {
			return
		}

		space += count

		printBST(node.Right, space)

		fmt.Print("\n")
		for i := count; i < space; i++ {
			fmt.Print(" ")
		}
		fmt.Println(node.Value)

		printBST(node.Left, space)
	}

	printBST(tree, 0)
}
