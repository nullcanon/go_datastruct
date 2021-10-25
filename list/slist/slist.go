package slist

import (
	"data_struct/utils"
)

/*

		   head					  tail
		 	|						|
			v						v
	      Node1 ----> Node2 ----> Node3 ----> nil

*/

type Node struct {
	next  *Node
	Value interface{}
}

func (n *Node) Next() *Node {
	return n.next
}

type List struct {
	head *Node
	tail *Node
	len  uint
}

func New() *List {
	list := &List{}
	return list
}

func (l *List) Len() uint {
	return l.len
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) PushBack(v interface{}) {
	n := &Node{Value: v}
	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.len++
}

func (l *List) PushFront(v interface{}) {
	n := &Node{Value: v}
	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		l.head.next = n
		l.head = n
	}
	l.len++
}

func (l *List) Insert(v interface{}, pos uint) {

	if pos > l.len {
		pos = l.len
	}

	n := &Node{Value: v}
	if pos == 0 {
		if l.len != 0 {
			n.next = l.head.next
		}
		l.head = n
		l.tail = n
	} else if pos == l.len {
		l.tail.next = n
		l.tail = n
	} else {
		for cur := l.head; ; cur = cur.Next() {
			pos--
			if pos == 0 {
				n.next = cur.next
				cur.next = n
				break
			}
		}
	}
	l.len++
}

func (l *List) Find(v interface{}) int {
	pos := -1
	for cur := l.head; cur != nil; cur = cur.Next() {
		pos++
		if cur.Value == v {
			return pos
		}
	}
	return -1
}

func (l *List) Remove(pos int) {
	if pos < 0 || l.len == 0 || pos > int(l.len) {
		return
	}
	if pos == 0 {
		l.head = l.head.next
	} else {
		for cur := l.head; ; cur = cur.Next() {
			if pos == 1 {
				if cur.next == l.tail {
					l.tail = cur
					break
				}
				cur.next = cur.next.next
				break
			}
			pos--
		}
	}
	l.len--
}

func (l *List) Empty() bool {
	return l.len == 0
}

func (l *List) Back() interface{} {
	return l.tail.Value
}

func (l *List) Front() interface{} {
	return l.head.Value
}

func (l *List) Traverse(visitor utils.Visitor) {
	for n := l.head; n != nil; n = n.Next() {
		visitor(n.Value)
	}
}
