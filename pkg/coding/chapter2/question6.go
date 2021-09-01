package chapter2

func IsPalindrome(list *SinglyLinkedList) bool {
	if list.head == nil {
		return false
	}
	store := make(map[int]int)
	currentNode := list.head
	index := 0
	for currentNode != nil {
		store[currentNode.value] += index
		index++
		currentNode = currentNode.next
	}

	length := index
	isOdd := false
	for _, indexSum := range store {
		if indexSum != length-1 {
			if !isOdd && indexSum == length/2 {
				isOdd = true
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// todo 本に乗ってた別の解法を実装してみる！　並び替えと比較　再帰
