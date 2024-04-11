package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Value == value {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}
}

func main() {
	ll := &LinkedList{}
	ll.Insert(5)
	ll.Insert(6)
	ll.Insert(7)
	ll.Print() // 5 -> 6 -> 7 -> nil

	ll.Delete(6)
	ll.Print() // 5 -> 7 -> nil
}
