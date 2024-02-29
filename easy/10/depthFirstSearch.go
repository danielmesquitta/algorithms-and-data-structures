package easy10

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, children := range n.Children {
		array = children.DepthFirstSearch(array)
	}
	return array
}
