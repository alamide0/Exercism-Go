package linkedlist

import (
	"errors"
)

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(ints []int) *List {

	list := new(List)

	for _, v := range ints {
		list.Push(v)
	}

	return list
}

func (list *List) Push(data int) {

	list.head = &Element{data, list.head}

	list.size++
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Array() []int {

	res := make([]int, list.size)

	count := list.size - 1
	for p := list.head; p != nil; p = p.next {
		res[count] = p.data
		count--
	}

	return res
}

func (list *List) Pop() (int, error) {
	if list.size == 0 {
		return 0, errors.New("no data")
	}

	res := list.head.data

	list.head = list.head.next
	list.size--

	return res, nil
}

func (list *List) Reverse() *List {

	array := list.Array()

	for i, end := 0, len(array)/2; i < end; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}

	return New(array)
}
