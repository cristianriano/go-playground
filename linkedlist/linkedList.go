package linkedlist

import "fmt"

type MyLinkedList struct {
	head *node
	size int
}

type node struct {
	val int
	next *node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (list *MyLinkedList) Get(index int) int {
	if index > list.size - 1 { return -1 }

	p := list.head
	val := p.val
	for i := 0; i < index; i++ {
		p = p.next
		val = p.val
	}

	return val
}

func (list *MyLinkedList) AddAtHead(val int) {
	defer func(){ list.size++ }()

	n := node{val: val}
	if list.head == nil {
		list.head = &n
		return
	}

	n.next = list.head
	list.head = &n
}


func (list *MyLinkedList) AddAtTail(val int) {
	defer func(){ list.size++ }()

	if list.head == nil {
		list.head = &node{val: val}
		return
	}

	p := list.head
	for ; p.next != nil; p = p.next {}
	p.next = &node{val: val}
}

func (list *MyLinkedList) AddAtIndex(index, val int) {
	if index > list.size { return }
	if index == 0 {
		list.AddAtHead(val)
		return
	}
	if index == list.size {
		list.AddAtTail(val)
		return
	}

	p, i := list.head, 0
	for ; i < index - 2; i++ {
		p = p.next
	}
	n := node{val: val, next: p.next}
	p.next = &n
	list.size++
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index > list.size - 1 { return }

	list.size--
	if index == 0 {
		list.head = list.head.next
		return
	}

	p, i := list.head, 0
	for ; i < index - 1 ; i++ {
		p = p.next
	}
	p.next = p.next.next
}

func (list *MyLinkedList) Print() {
	if list.head == nil {
		fmt.Println("()")
	} else {
		p := list.head

		for ; p.next != nil; p = p.next {
			fmt.Printf("(%v)->", p.val)
		}
		fmt.Printf("(%v)\n", p.val)
	}
}
