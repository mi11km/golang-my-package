package chapter2

import (
	"testing"
)

func CreateLoopList(vals []int, index int) *SinglyLinkedList {
	l := GetLinkedListFromValues(vals)
	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = l.getNode(index)
	return l
}

func TestFindLoopNodeNegative(t *testing.T) {
	cases := []struct {
		vals []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1},
		},
		{
			[]int{},
		},
	}
	for _, c := range cases {
		l := GetLinkedListFromValues(c.vals)
		actual := l.FindLoopNode()
		if actual != nil {
			t.Fatalf("Expected: nil, actual: %v\n", actual)
		}
	}
}

func TestFindLoopNodePositive(t *testing.T) {
	cases := []struct {
		vals  []int
		index int
	}{
		{
			[]int{100, 101, 102, 103, 104},
			0,
		},
		{
			[]int{100, 101, 102, 103, 104},
			2,
		},
		{
			[]int{100, 101, 102, 103, 104},
			4,
		},
	}
	for _, c := range cases {
		l := CreateLoopList(c.vals, c.index)
		actual := l.FindLoopNode()
		if actual.value != c.index+100 {
			t.Fatalf("Expected: %v, actual: %v\n", c.index, actual)
		}
	}
}
