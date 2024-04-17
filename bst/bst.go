package bst

import (
	"encoding/json"
	"fmt"
	"math"
)

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
	var validate func(*BST, int, int) bool
	validate = func(node *BST, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Value < min || node.Value >= max {
			return false
		}
		return validate(node.Left, min, node.Value) &&
			validate(node.Right, node.Value, max)
	}

	return validate(tree, math.MinInt, math.MaxInt)
}

func (tree *BST) InOrderTraverse() []int {
	return tree.inOrderTraverse(&[]int{})
}

func (tree *BST) inOrderTraverse(array *[]int) []int {
	if tree == nil {
		return *array
	}
	tree.Left.inOrderTraverse(array)
	*array = append(*array, tree.Value)
	tree.Right.inOrderTraverse(array)
	return *array
}

func (tree *BST) PreOrderTraverse() []int {
	return tree.preOrderTraverse(&[]int{})
}

func (tree *BST) preOrderTraverse(array *[]int) []int {
	if tree == nil {
		return *array
	}
	*array = append(*array, tree.Value)
	tree.Left.preOrderTraverse(array)
	tree.Right.preOrderTraverse(array)
	return *array
}

func (tree *BST) PostOrderTraverse() []int {
	return tree.postOrderTraverse(&[]int{})
}

func (tree *BST) postOrderTraverse(array *[]int) []int {
	if tree == nil {
		return *array
	}
	tree.Left.postOrderTraverse(array)
	tree.Right.postOrderTraverse(array)
	*array = append(*array, tree.Value)
	return *array
}

func (tree *BST) Print() {
	var print func(*BST, string, string)
	print = func(node *BST, prefix, childrenPrefix string) {
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

func (tree *BST) FromJSON(jsonTree JSONTree) {
	nodes := make(map[string]*BST)
	for _, node := range jsonTree.Tree.Nodes {
		nodes[node.ID] = &BST{Value: node.Value}
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

func (tree *BST) FromJSONString(jsonData string) {
	jsonTree := JSONTree{}
	err := json.Unmarshal([]byte(jsonData), &jsonTree)
	if err != nil {
		panic(err)
	}
	tree.FromJSON(jsonTree)
}
