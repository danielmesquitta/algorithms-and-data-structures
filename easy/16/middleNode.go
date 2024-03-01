package easy16

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) getDepth(count int) int {
	if ll == nil {
		return count
	}
	return ll.Next.getDepth(count + 1)
}

func (ll *LinkedList) getByDepth(count int) *LinkedList {
	if count == 1 {
		return ll
	}
	return ll.Next.getByDepth(count - 1)
}

func MiddleNode(ll *LinkedList) *LinkedList {
	totalDepth := ll.getDepth(0)

	depth := (totalDepth / 2) + 1

	return ll.getByDepth(depth)
}
