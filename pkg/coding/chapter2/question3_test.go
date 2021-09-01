package chapter2

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name string
		arg  int
		want []int
	}{
		{name: "successfully delete", arg: 2, want: []int{1, 2, 4, 5}},
		{name: "arg: nil", arg: -1, want: []int{1, 2, 3, 4, 5}},
		{name: "last node", arg: 4, want: []int{1, 2, 3, 4, 5}},
		{name: "first node", arg: 0, want: []int{2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, val := range data {
				l.Insert(val)
			}
			n := l.getNode(tt.arg)
			if DeleteNode(n) {
				l.length--
			}
			if got := l.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
