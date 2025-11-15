package linkedlist

import "errors"

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func NewList(elements ...interface{}) *List {
	if len(elements) == 0 {
		return new(List)
	}

	head := &Node{Value: elements[0], prev: nil}
	curr := head
	for _, v := range elements[1:] {
		curr.next = &Node{Value: v, prev: curr}
		curr = curr.next
	}

	return &List{
		head: head,
		tail: curr,
		size: len(elements),
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v, next: l.head}

	if l.size == 0 {
		l.tail = newNode
	} else {
		l.head.prev = newNode
	}

	l.head = newNode
	l.size++
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v, prev: l.tail}

	if l.size == 0 {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}

	l.tail = newNode
	l.size++
}

func (l *List) Shift() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("empty")
	}

	node := l.head
	l.head = node.next

	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}

	l.size--
	return node.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("empty")
	}

	node := l.tail
	l.tail = node.prev

	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	l.size--
	return node.Value, nil
}


func (l *List) Reverse() {
	var prev *Node
	curr := l.head
	for curr != nil {
		next := curr.next
		curr.prev = next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.tail = l.head
	l.head = prev
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
