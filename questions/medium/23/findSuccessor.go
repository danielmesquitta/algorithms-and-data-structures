package medium23

type BinaryTree struct {
	Value               int
	Left, Right, Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	inOrderTraversalNodes := []int{}
	successor := auxFindSuccessor(tree, node, &inOrderTraversalNodes)
	return successor
}

func auxFindSuccessor(
	tree *BinaryTree,
	node *BinaryTree,
	inOrderTraversalNodes *[]int,
) *BinaryTree {
	if tree.Left != nil {
		successor := auxFindSuccessor(tree.Left, node, inOrderTraversalNodes)
		if successor != nil {
			return successor
		}
	}

	*inOrderTraversalNodes = append(*inOrderTraversalNodes, tree.Value)
	index := len(*inOrderTraversalNodes) - 1

	if index > 0 {
		predecessor := (*inOrderTraversalNodes)[index-1]
		if isSuccessor := predecessor == node.Value; isSuccessor {
			return tree
		}
	}

	if tree.Right != nil {
		successor := auxFindSuccessor(tree.Right, node, inOrderTraversalNodes)
		if successor != nil {
			return successor
		}
	}

	return nil
}
