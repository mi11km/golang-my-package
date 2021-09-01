package chapter2

// AddListsByLoop リストで表された2数の和(下位の桁から順番に並んでる)
func AddListsByLoop(l1, l2 *SinglyLinkedList) *SinglyLinkedList {
	add := NewSinglyLinkedList()
	n1, n2 := l1.head, l2.head
	v1, v2, carryUp, remainder := 0, 0, 0, 0
	for n1 != nil || n2 != nil {
		if n1 != nil {
			v1 = n1.value
			n1 = n1.next
		} else {
			v1 = 0
		}
		if n2 != nil {
			v2 = n2.value
			n2 = n2.next
		} else {
			v2 = 0
		}
		remainder = (v1 + v2 + carryUp) % 10
		carryUp = (v1 + v2 + carryUp - remainder) / 10
		add.Insert(remainder)
	}
	return add
}

func AddLists(l1, l2 *SinglyLinkedList) *SinglyLinkedList {
	list := NewSinglyLinkedList()
	list.head = addListsByRecursion(l1.head, l2.head, 0)
	for n := list.head; n != nil; n = n.next {
		list.length++
	}
	return list
}

func addListsByRecursion(l1, l2 *node, carry int) *node {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	value := carry
	if l1 != nil {
		value += l1.value
	}
	if l2 != nil {
		value += l2.value
	}

	result := &node{}
	result.value = value % 10

	if l1 != nil || l2 != nil {
		n1, n2, v := &node{}, &node{}, 0
		if l1 == nil {
			n1 = nil
		} else {
			n1 = l1.next
		}
		if l2 == nil {
			n2 = nil
		} else {
			n2 = l2.next
		}
		if value >= 10 {
			v = 1
		} else {
			v = 0
		}
		more := addListsByRecursion(n1, n2, v)
		result.next = more
	}
	return result
}

/*
todo リストで表された2数の和(上位の桁から順番に並んでる)
*/
//func AddLinkedLists(l1, l2 *SinglyLinkedList) *SinglyLinkedList {
//
//}
