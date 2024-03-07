package easy16

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) countDepth(count int) int {
	if ll == nil {
		return count
	}
	return ll.Next.countDepth(count + 1)
}

func (ll *LinkedList) getNodeByDepth(depth int) *LinkedList {
	if depth == 1 {
		return ll
	}
	return ll.Next.getNodeByDepth(depth - 1)
}

func (ll *LinkedList) GetMiddleNode() *LinkedList {
	totalDepth := ll.countDepth(0)
	depth := (totalDepth / 2) + 1
	return ll.getNodeByDepth(depth)
}
