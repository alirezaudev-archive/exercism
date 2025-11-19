package linkedlist

import "errors"

type Node struct {
	val  int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func New(elements []int) *List {
	l := &List{size: len(elements)}
	if l.size == 0 {
		return l
	}

	curr := &Node{val: elements[0]}
	l.head = curr
	for _, v := range elements[1:] {
		curr.next = &Node{val: v}
		curr.next.prev = curr
		curr = curr.next
	}
	l.tail = curr

	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	node := &Node{val: element}
	if l.size == 0 {
		l.head = node
	} else {
		node.prev = l.tail
		l.tail.next = node
	}
	l.tail = node
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return -1, errors.New("empty")
	}

	popped := l.tail.val
	if l.size == 1 {
		l.head = nil
	}

	l.tail = l.tail.prev
	l.size--
	return popped, nil
}

func (l *List) Array() []int {
	elements := make([]int, l.size)
	curr := l.head
	for i := 0; i < l.size; i++ {
		elements[i] = curr.val
		curr = curr.next
	}

	return elements
}

func (l *List) Reverse() *List {
	var prev *Node

	l.tail = l.head
	curr := l.head
	for curr != nil {
		next := curr.next
		curr.next, curr.prev = prev, next
		prev = curr
		curr = next
	}

	l.head = prev
	return l
}
