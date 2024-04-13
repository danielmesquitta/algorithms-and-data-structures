package easy10

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(arr []string) []string {
	arr = append(arr, n.Name)
	for _, child := range n.Children {
		arr = child.DepthFirstSearch(arr)
	}
	return arr
}
