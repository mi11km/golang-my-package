package chapter2

// FindLoopNode 循環する連結リストの循環点を返す関数
func (l *SinglyLinkedList) FindLoopNode() *node {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return nil
	}

	slow = l.head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return fast
}
