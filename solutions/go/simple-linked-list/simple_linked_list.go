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
	if l.root == nil {
		l.root = &Element{element, nil}
	} else {
		var cur *Element
		for cur = l.root; cur.next != nil; cur = cur.next {
		}
		cur.next = &Element{element, nil}
	}
}

func (l *List) Pop() (int, error) {
	if l.root == nil {
		return 0, errors.New("cannot pop from empty list")
	}
	var val int
	if l.root.next == nil {
		val = l.root.value
		l.root = nil
	} else {
		var cur *Element
		for cur = l.root; cur.next.next != nil; cur = cur.next {
		}
		val = cur.next.value
		cur.next = nil
	}
	return val, nil
}

func (l *List) Array() []int {
	var arr []int
	for cur := l.root; cur != nil; cur = cur.next {
		arr = append(arr, cur.value)
	}
	return arr
}

func (l *List) Reverse() *List {
	newList := &List{}
	for cur := l.root; cur != nil; cur = cur.next {
		if newList.root == nil {
			newList.Push(cur.value)
		} else {
			newList.root = &Element{cur.value, newList.root}
		}
	}
	return newList

}
