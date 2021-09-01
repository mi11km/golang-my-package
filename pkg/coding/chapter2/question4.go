package chapter2

// SplitList 連結リストの要素を並び替えて、xより小さいものが前に来るようにするメソッド
//           計算量 O(n^2) で遅い
func (l *SinglyLinkedList) SplitList(x int) {
	node1 := l.head
	for i := 0; i < l.length; i++ {
		if node1.value >= x {
			node2 := node1.next
			for j := i + 1; j < l.length; j++ {
				if node2.value < x {
					l.Swap(i, j)
					break
				}
				node2 = node2.next
			}
		}
		node1 = node1.next
	}
}

func (l *SinglyLinkedList) Swap(i1, i2 int) bool {
	if !(0 <= i1 && i1 < l.length) || !(0 <= i2 && i2 < l.length) {
		return false
	}
	n1, n2 := l.head, l.head
	for i := 0; i < i1; i++ {
		n1 = n1.next
	}
	for i := 0; i < i2; i++ {
		n2 = n2.next
	}
	n1.value, n2.value = n2.value, n1.value
	return true
}

// Partition 存在するノードを用いて、前方に新しいリストを作りながら作成する。最後に旧リストを削除
//			 xよりより小さいものは先頭に挿入し、それ以外は末尾（新リストと旧リストの間）に挿入する。
func (l *SinglyLinkedList) Partition(x int) {
	n := l.head
	head := l.head
	tail := l.head
	for n != nil {
		next := n.next
		if n.value < x {
			n.next = head
			head = n
		} else {
			tail.next = n
			tail = n
		}
		n = next
	}
	tail.next = nil
	l.head = head
}

/*
新しくxより小さいリストと, x以上のリストを作成しそれを連結する（順番はそのまま）
*/
func (l *SinglyLinkedList) Partition2(x int) {
	var beforeStart, beforeEnd, afterStart, afterEnd, n *node = nil, nil, nil, nil, l.head
	for n != nil {
		next := n.next
		n.next = nil
		if n.value < x {
			// 前半のリストの最後のノードに挿入する
			if beforeStart == nil {
				beforeStart = n
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = n
				beforeEnd = n
			}
		} else {
			// 後半のリストの最後にノードを挿入する
			if afterStart == nil {
				afterStart = n
				afterEnd = afterStart
			} else {
				afterEnd.next = n
				afterEnd = n
			}
		}
		n = next
	}
	if beforeStart == nil {
		l.head = afterStart
	} else {
		beforeEnd.next = afterStart
		l.head = beforeStart
	}
}
