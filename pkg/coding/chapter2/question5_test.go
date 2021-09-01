package chapter2

import (
	"reflect"
	"testing"
)

func TestAddListsByLoop(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"success", args{[]int{7, 1, 6}, []int{5, 9, 2}}, []int{2, 1, 9}},
		{"digits are different (larger l1)", args{[]int{7, 1, 6, 6}, []int{5, 9, 2}}, []int{2, 1, 9, 6}},
		{"digits are different (larger l2)", args{[]int{7, 1, 6}, []int{5, 1, 7, 8}}, []int{2, 3, 3, 9}},
		{"both empty", args{[]int{}, []int{}}, []int{}},
		{"one side empty ", args{[]int{0}, []int{}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1, list2 := NewSinglyLinkedList(), NewSinglyLinkedList()
			for _, v := range tt.args.l1 {
				list1.Insert(v)
			}
			for _, v := range tt.args.l2 {
				list2.Insert(v)
			}

			if got := AddListsByLoop(list1, list2).Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}

func TestAddLists(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"success", args{[]int{7, 1, 6}, []int{5, 9, 2}}, []int{2, 1, 9}},
		{"digits are different (larger l1)", args{[]int{7, 1, 6, 6}, []int{5, 9, 2}}, []int{2, 1, 9, 6}},
		{"digits are different (larger l2)", args{[]int{7, 1, 6}, []int{5, 1, 7, 8}}, []int{2, 3, 3, 9}},
		{"both empty", args{[]int{}, []int{}}, []int{}},
		{"one side empty ", args{[]int{0}, []int{}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1, list2 := NewSinglyLinkedList(), NewSinglyLinkedList()
			for _, v := range tt.args.l1 {
				list1.Insert(v)
			}
			for _, v := range tt.args.l2 {
				list2.Insert(v)
			}

			if got := AddLists(list1, list2).Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
