package linkedlist

import "errors"

type Element struct {
	value int
	next  *Element
}

type List struct {
	root *Element
}

func New(elements []int) *List {
	l := &List{}
	for _, v := range elements {
		l.Push(v)
	}
	return l
}

func (l *List) Size() int {
	var i int
	var cur *Element
	for i, cur = 0, l.root; cur != nil; i, cur = i+1, cur.next {
	}
	return i
}

func (l *List) Push(element int) {
	l.root = &Element{element, l.root}
}

func (l *List) Pop() (int, error) {
	if l.root == nil {
		return 0, errors.New("cannot pop from empty list")
	}
	val := l.root.value
	l.root = l.root.next
	return val, nil
}

func (l *List) Array() []int {
	size := l.Size()
	arr := make([]int, size)
	for i, cur := size-1, l.root; cur != nil; i, cur = i-1, cur.next {
		arr[i] = cur.value
	}
	return arr
}

func (l *List) Reverse() *List {
	newList := &List{}
	for cur := l.root; cur != nil; cur = cur.next {
		newList.Push(cur.value)
	}
	return newList

}
