package chapter2

func FindIntersection(l1, l2 *SinglyLinkedList) *node {
	node1 := l1.head
	node2 := l2.head

	/* 共通ノードがある → 交わってる → 途中で合流してあと一緒 → 合流点から最後のノードまでが一致 */
	if node1 == nil || node2 == nil {
		return nil
	}

	if getTail(l1) != getTail(l2) {
		return nil
	}

	if l1.length > l2.length {
		for i := 0; i < l1.length-l2.length; i++ {
			node1 = node1.next
		}
	} else {
		for i := 0; i < l2.length-l1.length; i++ {
			node2 = node2.next
		}
	}

	for node1 != node2 {
		node1 = node1.next
		node2 = node2.next
	}

	return node1
}

func getTail(list *SinglyLinkedList) *node {
	if list.head == nil {
		return nil
	}
	tail := list.head
	for tail.next != nil {
		tail = tail.next
	}
	return tail
}

/* ２重ループパターン */
//for node1 != nil {
//	for node2 != nil {
//		if node1 == node2 {
//			return node1
//		}
//		node2 = node2.next
//	}
//	node1 = node1.next
//	node2 = l2.head
//}

/* ハッシュテーブル使う(円環になってない前提) */
//nodeMap := make(map[*node]int)
//for node1 != nil {
//	nodeMap[node1]++
//	node1 = node1.next
//}
//for node2 != nil {
//	nodeMap[node2]++
//	node2 = node2.next
//}
//for commonNode, val := range nodeMap {
//	if val >= 2 {
//		return commonNode
//	}
//}
