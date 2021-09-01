package chapter2

// DeleteDuplicateVal ソートされていない連結リストから、重複する要素を削除する
func (ll *DoublyLinkedList) DeleteDuplicateVal() {
	l := ll.length
	for i := 0; i < l; i++ {
		val := ll.Get(i)
		for j := i + 1; j < l; {
			if val == ll.Get(j) {
				ll.Delete(j)
				l--
			} else {
				j++
			}
		}
	}
}

func (ll *DoublyLinkedList) DeleteDups() {
	seen := make(map[int]struct{})
	current := ll.head
	deleteIndex := 0
	for current != nil {
		if _, ok := seen[current.value]; ok {
			ll.Delete(deleteIndex)
		} else {
			seen[current.value] = struct{}{}
			deleteIndex++
		}
		current = current.next
	}
}
