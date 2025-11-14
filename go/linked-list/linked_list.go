package linkedlist

import "errors"

type Node struct {
	Value any
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...any) *List {
	l := &List{}
	for _, elem := range elements {
		l.Push(elem)
	}
	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

// Unshift inserts a node in the beginning (head)
func (l *List) Unshift(v any) {
	n := &Node{
		Value: v,
		prev:  nil,
		next:  l.head,
	}
	if l.head == nil {
		l.tail = n
	} else {
		l.head.prev = n
	}
	l.head = n
}

// Push inserts a node at the end (tail)
func (l *List) Push(v any) {
	n := &Node{
		Value: v,
		prev:  l.tail,
		next:  nil,
	}
	if l.tail == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

// Shift removes a node from the beginning (head)
func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	n := l.head
	l.head = n.next
	if l.head == nil { // This was the last node
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return n.Value, nil
}

// Pop removes a node from the end (tail)
func (l *List) Pop() (any, error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	n := l.tail
	l.tail = n.prev
	if l.tail == nil { // This was the last node
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return n.Value, nil
}

func (l *List) Reverse() {
	for n := l.head; n != nil; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
