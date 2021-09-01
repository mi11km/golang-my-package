package chapter2

import "testing"

func TestSinglyLinkedList_GetFromRear(t *testing.T) {

	type args struct {
		index int
		list  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{index: 1, list: []int{1, 2, 3, 4}}, want: 4},
		{name: "success", args: args{index: 2, list: []int{1, 2, 3, 4}}, want: 3},
		{name: "success", args: args{index: 3, list: []int{1, 2, 3, 4}}, want: 2},
		{name: "success", args: args{index: 4, list: []int{1, 2, 3, 4}}, want: 1},
		{name: "index out of range", args: args{index: 5, list: []int{1, 2, 3, 4}}, want: -1},
		{name: "index out of range", args: args{index: -1, list: []int{1, 2, 3, 4}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, val := range tt.args.list {
				l.Insert(val)
			}
			if got := l.GetFromRear(tt.args.index); got != tt.want {
				t.Errorf("GetFromRear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_KthToLast(t *testing.T) {

	type args struct {
		index int
		list  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{index: 1, list: []int{1, 2, 3, 4}}, want: 4},
		{name: "success", args: args{index: 2, list: []int{1, 2, 3, 4}}, want: 3},
		{name: "success", args: args{index: 3, list: []int{1, 2, 3, 4}}, want: 2},
		{name: "success", args: args{index: 4, list: []int{1, 2, 3, 4}}, want: 1},
		{name: "index out of range", args: args{index: 5, list: []int{1, 2, 3, 4}}, want: -1},
		{name: "index out of range", args: args{index: -1, list: []int{1, 2, 3, 4}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, val := range tt.args.list {
				l.Insert(val)
			}
			if got := l.KthToLast(tt.args.index); got != tt.want {
				t.Errorf("KthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
