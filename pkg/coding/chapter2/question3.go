package chapter2

// DeleteNode 単結合リストにおいて、間の要素で、その要素のみにアクセス可能の場合、その要素を削除する関数
func DeleteNode(n *node) bool {
	if n == nil || n.next == nil {
		return false
	}
	next := n.next
	n.value = next.value
	n.next = next.next
	return true
}
