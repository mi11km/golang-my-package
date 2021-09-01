package chapter2

import (
	"strconv"
	"strings"
)

type DoublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	value int
	prev  *node
	next  *node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

// Insert at the end
func (ll *DoublyLinkedList) Insert(val int) {
	newNode := &node{value: val}
	ll.length++
	if ll.tail == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
}

func (ll *DoublyLinkedList) Delete(index int) {
	if ll.length <= index {
		return
	}
	if index == 0 {
		if ll.length == 1 {
			ll.head, ll.tail, ll.length = nil, nil, 0
		} else {
			ll.head = ll.head.next
			ll.head.prev = nil
			ll.length--
		}
	} else if ll.length == index+1 {
		ll.tail = ll.tail.prev
		ll.tail.next = nil
		ll.length--
	} else {
		deleteNode := ll.head
		for i := 0; i < index; i++ {
			deleteNode = deleteNode.next
		}
		deleteNode.next.prev = deleteNode.prev
		deleteNode.prev.next = deleteNode.next
		ll.length--
	}
}

func (ll *DoublyLinkedList) getNode(index int) *node {
	if ll.length <= index {
		return nil
	}
	n := ll.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n
}

func (ll *DoublyLinkedList) Get(index int) int {
	if n := ll.getNode(index); n != nil {
		return n.value
	}
	return -1
}

func (ll *DoublyLinkedList) Len() int {
	return ll.length
}

/*--------------*/

// Mostly used for testing and debug

func (ll *DoublyLinkedList) Slice() []int {
	slice := make([]int, ll.length)
	n := ll.head
	for i := 0; i < ll.length; i++ {
		slice[i] = n.value
		n = n.next
	}
	return slice
}

func (ll *DoublyLinkedList) SliceFromTail() []int {
	slice := make([]int, ll.length)
	n := ll.tail
	for i := 0; i < ll.length; i++ {
		slice[i] = n.value
		n = n.prev
	}
	return slice
}

func (ll *DoublyLinkedList) String() string {
	n := ll.head
	s := strings.Builder{}
	s.WriteByte(byte('['))
	for n != nil {
		s.WriteString(strconv.Itoa(n.value))
		if n.next != nil {
			s.WriteByte(byte(' '))
		}
		n = n.next
	}
	s.WriteByte(byte(']'))
	return s.String()
}
