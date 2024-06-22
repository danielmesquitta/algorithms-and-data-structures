package medium25

import (
	"encoding/json"
	"fmt"
)

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MergeBinaryTrees(tree1 *BinaryTree, tree2 *BinaryTree) *BinaryTree {
	mergedTree := new(BinaryTree)
	auxMergeBinaryTrees(tree1, tree2, mergedTree)
	return mergedTree
}

func auxMergeBinaryTrees(
	tree1,
	tree2,
	mergedTree *BinaryTree,
) {
	if tree1 == nil {
		tree1 = new(BinaryTree)
	}
	if tree2 == nil {
		tree2 = new(BinaryTree)
	}

	mergedTree.Value = tree1.Value + tree2.Value

	if tree1.Left != nil || tree2.Left != nil {
		mergedTree.Left = new(BinaryTree)
		auxMergeBinaryTrees(tree1.Left, tree2.Left, mergedTree.Left)
	}

	if tree1.Right != nil || tree2.Right != nil {
		mergedTree.Right = new(BinaryTree)
		auxMergeBinaryTrees(tree1.Right, tree2.Right, mergedTree.Right)
	}
}

func (tree *BinaryTree) Print() {
	var print func(*BinaryTree, string, string)
	print = func(node *BinaryTree, prefix, childrenPrefix string) {
		if node != nil {
			fmt.Println(prefix + fmt.Sprint(node.Value))

			if node.Left != nil || node.Right != nil {
				if node.Left != nil {
					print(
						node.Left,
						childrenPrefix+"├── ",
						childrenPrefix+"│   ",
					)
				} else {
					fmt.Println(childrenPrefix + "├── ")
				}

				if node.Right != nil {
					print(
						node.Right,
						childrenPrefix+"└── ",
						childrenPrefix+"    ",
					)
				} else {
					fmt.Println(childrenPrefix + "└── ")
				}
			}
		}
	}
	print(tree, "", "")
}

type JSONTree struct {
	Tree Tree `json:"tree"`
}

type Tree struct {
	Nodes []Node `json:"nodes"`
	Root  string `json:"root"`
}

type Node struct {
	ID    string `json:"id"`
	Left  string `json:"left"`
	Right string `json:"right"`
	Value int    `json:"value"`
}

func (tree *BinaryTree) FromJSON(jsonTree JSONTree) {
	nodes := make(map[string]*BinaryTree)
	for _, node := range jsonTree.Tree.Nodes {
		nodes[node.ID] = &BinaryTree{Value: node.Value}
	}
	for _, node := range jsonTree.Tree.Nodes {
		if node.Left != "" {
			nodes[node.ID].Left = nodes[node.Left]
		}
		if node.Right != "" {
			nodes[node.ID].Right = nodes[node.Right]
		}
	}
	*tree = *nodes[jsonTree.Tree.Root]
}

func (tree *BinaryTree) FromJSONString(jsonData string) {
	jsonTree := JSONTree{}
	err := json.Unmarshal([]byte(jsonData), &jsonTree)
	if err != nil {
		panic(err)
	}
	tree.FromJSON(jsonTree)
}
