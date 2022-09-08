package main

import "fmt"

//Linked List
type LinkedList struct {
	head   *Node
	length int
}
type Node struct {
	key  int
	next *Node
}

func (l *LinkedList) Append(Val int) {
	if l.head == nil {
		l.head = &Node{key: Val}
		l.length++
	} else {
		ins := l.head
		for ins.next != nil {
			ins = ins.next
		}
		ins.next = &Node{key: Val}
		l.length++
	}
}
func (l *LinkedList) Traverse() {
	if l.head == nil {
		return
	} else {
		toPrint := l.head
		for toPrint != nil {
			fmt.Println(toPrint.key)
			toPrint = toPrint.next
			l.length--
		}
	}
}
func (l *LinkedList) Prepend(val int) {
	if l.head == nil {
		l.head = &Node{key: val}
	} else {
		second := l.head
		l.head = &Node{key: val, next: second}
	}
}
func (l *LinkedList) MidPend(val int, new int) {
	if l.head == nil {
		l.head = &Node{key: val}
	} else {
		ptr := l.head
		for ptr.next != nil {
			if ptr.key == val {
				ptr.next = &Node{key: new, next: ptr.next}
			}
			ptr = ptr.next
		}

	}
}

func main() {
	l := &LinkedList{}
	l.Append(9)
	l.Append(10)
	l.Append(11)
	l.Append(11)
	l.Prepend(21)
	l.MidPend(21, 22)
	l.Traverse()
}
