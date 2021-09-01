package chapter2

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	type args struct {
		l1     []int
		l2     []int
		common []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{l1: []int{1, 2, 3}, l2: []int{4, 5, 6}, common: []int{9, 8, 7}}},
		{"success diff len", args{l1: []int{1, 2, 3, 4}, l2: []int{4, 5, 6}, common: []int{9, 8, 7}}},
		{"success diff len", args{l1: []int{1, 2, 3}, l2: []int{4, 5, 6, 7}, common: []int{9, 8, 7}}},
		{"success nil", args{l1: []int{}, l2: []int{4, 5, 6, 7}, common: []int{9, 8, 7}}},
		{"success", args{l1: []int{1, 2, 3}, l2: []int{4, 5, 6, 7, 8}, common: []int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common := NewSinglyLinkedList()
			for _, v := range tt.args.common {
				common.Insert(v)
			}
			list1 := NewSinglyLinkedList()
			for _, v := range tt.args.l1 {
				list1.Insert(v)
			}
			list2 := NewSinglyLinkedList()
			for _, v := range tt.args.l2 {
				list2.Insert(v)
			}
			if list1.head != nil && list2.head != nil {
				l1Tail := list1.getNode(len(tt.args.l1) - 1)
				l2Tail := list2.getNode(len(tt.args.l2) - 1)
				l1Tail.next = common.head
				l2Tail.next = common.head
			} else {
				common.head = nil
			}

			if got := FindIntersection(list1, list2); !reflect.DeepEqual(got, common.head) {
				t.Errorf("FindIntersection() = %v, want %v", got, common.head)
			}
		})
	}
}
