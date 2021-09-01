package chapter2

func (l *SinglyLinkedList) GetFromRear(index int) int {
	if index < 0 || l.length < index {
		return -1
	}
	n := l.head
	for i := 0; i < l.length-index; i++ {
		n = n.next
	}
	return n.value
}

// KthToLast 長さがわかってない場合
func (l *SinglyLinkedList) KthToLast(index int) int {
	if index < 0 {
		return -1
	}
	p1 := l.head
	p2 := l.head
	for i := 0; i < index; i++ {
		if p1 == nil {
			return -1
		}
		p1 = p1.next
	}

	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2.value
}
