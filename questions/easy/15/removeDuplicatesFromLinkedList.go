package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) RemoveDuplicates() {
	node := ll
	for node != nil {
		if node.Next != nil && node.Value == node.Next.Value {
			*node = *node.Next
			continue
		}
		node = node.Next
	}
}
