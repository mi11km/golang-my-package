package chapter2

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList_DeleteDuplicateVal(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "success",
			arg:  []int{1, 2, 3, 3, 3, 3, 4, 5, 6, 7, 7, 8, 9},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "empty",
			arg:  []int{},
			want: []int{},
		},
		{
			name: "not change",
			arg:  []int{1},
			want: []int{1},
		},
		{
			name: "success",
			arg:  []int{3, 5, 2, 6, 1, 2, 3, 2, 5, 6, 2},
			want: []int{3, 5, 2, 6, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewDoublyLinkedList()
			for _, val := range tt.arg {
				ll.Insert(val)
			}

			ll.DeleteDuplicateVal()
			if !reflect.DeepEqual(ll.Slice(), tt.want) {
				t.Errorf("got: %v, want: %v\n", ll.String(), tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_DeleteDups(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "success",
			arg:  []int{1, 2, 3, 3, 3, 3, 4, 5, 6, 7, 7, 8, 9},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "empty",
			arg:  []int{},
			want: []int{},
		},
		{
			name: "empty",
			arg:  []int{1},
			want: []int{1},
		},
		{
			name: "success",
			arg:  []int{3, 5, 2, 6, 1, 2, 3, 2, 5, 6, 2},
			want: []int{3, 5, 2, 6, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewDoublyLinkedList()
			for _, val := range tt.arg {
				ll.Insert(val)
			}

			ll.DeleteDups()
			if !reflect.DeepEqual(ll.Slice(), tt.want) {
				t.Errorf("got: %v, want: %v\n", ll.String(), tt.want)
			}
		})
	}
}
