package linkedlist

import "errors"

type List struct {
	head *Node
	size int
}

type Node struct {
	element int
	next    *Node
}

func New(elements []int) *List {
	l := &List{}
	for _, elem := range elements {
		l.Push(elem)
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.size++
	n := &Node{element: element}
	if l.head == nil {
		l.head = n
		return
	}
	var tail *Node
	for tail = l.head; tail.next != nil; tail = tail.next {
	}
	tail.next = n
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("empty list")
	}
	l.size--
	if l.head.next == nil {
		element := l.head.element
		l.head = nil
		return element, nil
	}
	var newTail *Node
	for newTail = l.head; newTail.next != nil && newTail.next.next != nil; newTail = newTail.next {
	}
	tail := newTail.next
	newTail.next = nil
	return tail.element, nil
}

func (l *List) Array() []int {
	if l == nil {
		return []int{}
	}
	var elements []int
	for p := l.head; p != nil; p = p.next {
		elements = append(elements, p.element)
	}
	return elements
}

func (l *List) Reverse() *List {
	rl := &List{}
	elements := l.Array()
	for i := len(elements) - 1; i >= 0; i-- {
		rl.Push(elements[i])
	}
	return rl
}
