package easy15

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) SelfDelete() {
	if ll.Next != nil {
		*ll = *ll.Next
	}
}

func (ll *LinkedList) RemoveDuplicates() {
	if ll.Next == nil {
		return
	}
	if ll.Value == ll.Next.Value {
		ll.SelfDelete()
		ll.RemoveDuplicates()
		return
	}
	ll.Next.RemoveDuplicates()
}

func (ll *LinkedList) RemoveDuplicatesFromLinkedList() *LinkedList {
	if ll == nil || ll.Next == nil {
		return ll
	}
	ll.Next = ll.Next.RemoveDuplicatesFromLinkedList()
	if ll.Value == ll.Next.Value {
		return ll.Next
	}
	return ll
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	ll := linkedList
	for ll != nil {
		if ll.Next != nil && ll.Value == ll.Next.Value {
			*ll = *ll.Next
			continue
		}
		ll = ll.Next
	}
	return linkedList
}
