package linkedlist

import (
	"errors"
)

var ErrEmptyList = errors.New("empty list")

type Node struct {
	Val            interface{}
	previous, next *Node
}

type List struct {
	head, tail *Node
	size       int
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

func NewList(args ...interface{}) *List {
	list := &List{}
	for _, v := range args {
		list.PushBack(v)
	}
	return list
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) PushFront(v interface{}) {
	node := &Node{v, nil, nil}

	if l.size == 0 {
		l.tail = node
	} else {
		node.next = l.head
		l.head.previous = node
	}

	l.head = node

	l.size++
}

func (l *List) PushBack(v interface{}) {
	node := &Node{v, nil, nil}
	if l.size == 0 {
		l.head = node
	} else {
		node.previous = l.tail
		l.tail.next = node
	}

	l.tail = node

	l.size++
}

func (l *List) PopFront() (interface{}, error) {
	if l.size == 0 {
		return 0, ErrEmptyList
	}
	node := l.head

	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = node.next
		node.next.previous = nil
		node.next = nil
	}

	l.size--
	return node.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.size == 0 {
		return 0, ErrEmptyList
	}
	node := l.tail
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = node.previous
		node.previous.next = nil
		node.previous = nil
	}

	l.size--
	return node.Val, nil
}

func (l *List) Reverse() *List {

	h, t := l.head, l.tail

	for i, end := 0, l.size/2; i < end; i++ { //just reverse val
		h.Val, t.Val = t.Val, h.Val
		h = h.next
		t = t.previous
	}

	return l
}
